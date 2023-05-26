// @Author: Ciusyan 2023/2/7
package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"github.com/Go-To-Byte/DouSheng/dou-kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou-kit/exception"
	"github.com/Go-To-Byte/DouSheng/dou-kit/exception/custom"
	"github.com/Go-To-Byte/DouSheng/video-service/apps/video"

	"github.com/Go-To-Byte/DouSheng/api-rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/api-rooter/common/utils"
)

type feedResp struct {
	*custom.CodeMsg
	*video.FeedSetResponse
}

type listResp struct {
	*custom.CodeMsg
	*video.PublishListResponse
}

func (h *Handler) publishAction(ctx *gin.Context) error {

	// 1、构建 GRPC 请求参数
	req := video.NewPublishVideoRequest()
	if err := ctx.Bind(req); err != nil {
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}
	// 从认证成功后的Token中，取出传递下来的Token对象，// TODO 将其放在GRPC的 metadata 中
	if value, exists := ctx.Get(constant.REQUEST_TOKEN); exists {
		// 从Token中 -> 给UserId赋值
		req.UserId = value.(*token.Token).GetUserId()
	}

	// 2、读取文件数据并且上传
	file, err := ctx.FormFile(constant.REQUEST_FILE)
	if err != nil {
		return exception.WithStatusCode(constant.BAD_NO_FILE)
	}

	// 3、文件上传
	// 	用于传递文件上传的结果
	uploadMsgCh := make(chan *utils.UploadMsg)

	// 	防止上传太久，都没有发信号出来。
	uploadCtx, uploadCancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer uploadCancel()
	// 	这里只管上传视频，会通过 uploadMsgCh 发送信号出来
	go utils.UploadFile(uploadCtx, file, uploadMsgCh)

	// 4、将信息入库
	publishErrCh := make(chan error)
	go func(publishErrCh chan error) {

		// 这里需要等待上传的结果
		select {
		case <-uploadCtx.Done():
			// 说明上传超时了，没必要等了，直接给一个上传上传失败的错误
			publishErrCh <- exception.WithStatusCode(constant.BAD_UPLOAD_FILE)
			h.l.Errorf("video: publish 超时了 %s", uploadCtx.Err())
			return
		case uploadMsg := <-uploadMsgCh:
			if uploadMsg.Err != nil {
				// 说明上传失败了
				h.l.Errorf("video: publish 文件上传错误 %s", uploadMsg.Err)
				publishErrCh <- exception.WithStatusCode(constant.BAD_UPLOAD_FILE)
				return
			}

			// 文件上传成功后，给文件参数赋值
			req.CoverUrl = uploadMsg.UploadResult.CoverRelativeURI
			req.PlayUrl = uploadMsg.UploadResult.RelativeURI

			_, err = h.service.PublishVideo(ctx.Request.Context(), req)
			if err != nil {
				// 来到这里，说明入库失败了吗，需要将已上传的视频 Delete
				publishErrCh <- exception.GrpcErrWrapper(err)

				// 这里面本身就会在后台删除，所以不用再起一个协程了
				uploadMsg.Delete()
				return
			}

			// 来到这里说明保存成功了
			publishErrCh <- nil
		}

	}(publishErrCh)

	// 阻塞监听
	select {
	case err := <-publishErrCh:
		if err != nil {
			// 说明上传 or 入库失败了
			return err
		}

		// 说明成功了
		ctx.JSON(http.StatusOK, custom.NewWithCode(constant.OPERATE_OK))
		return nil
	}
}

func (h *Handler) publishList(ctx *gin.Context) error {
	req := video.NewPublishListRequest()
	// 1、接收参数
	if err := ctx.ShouldBindQuery(req); err != nil {
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 业务请求
	videos, err := h.service.PublishList(ctx, req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	// 获取成功
	ctx.JSON(http.StatusOK, listResp{
		CodeMsg:             custom.Ok(constant.ACQUIRE_OK),
		PublishListResponse: videos,
	})
	return nil
}

func (h *Handler) feed(ctx *gin.Context) error {
	req := video.NewFeedVideosRequest()
	// 1、接收参数
	if err := ctx.ShouldBindQuery(req); err != nil {
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 业务请求
	videos, err := h.service.FeedVideos(ctx, req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	// 获取成功
	ctx.JSON(http.StatusOK, feedResp{
		CodeMsg:         custom.Ok(constant.ACQUIRE_OK),
		FeedSetResponse: videos,
	})
	return nil
}
