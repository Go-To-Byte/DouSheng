// @Author: Ciusyan 2023/2/7
package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"

	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception/custom"
	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
	"github.com/Go-To-Byte/DouSheng/video_service/common/utils"
)

type feedResp struct {
	*custom.CodeMsg
	*video.FeedSetResponse
}

type listResp struct {
	*custom.CodeMsg
	*video.PublishListResponse
}

func (h *Handler) publishAction(c context.Context, ctx *app.RequestContext) error {

	// 1、读取文件数据并且上传
	file, err := ctx.FormFile(constant.REQUEST_FILE)
	if err != nil {
		return exception.WithStatusCode(constant.BAD_NO_FILE)
	}

	uploaded, err := utils.UploadFile(file)
	if err != nil {
		return exception.WithStatusCode(constant.BAD_UPLOAD_FILE)
	}

	// TODO：若在下面发生错误，需要将已上传的视频 Delete
	req := video.NewPublishVideoRequest()
	// 2、接收参数
	if err := ctx.BindAndValidate(req); err != nil {
		h.log.Error(err)
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 文件上传成功后，给文件参数赋值
	req.CoverUrl = uploaded.CoverRelativeURI
	req.PlayUrl = uploaded.RelativeURI

	// 从认证成功后的Token中，取出传递下来的Token对象，
	if value, exists := ctx.Get(constant.REQUEST_TOKEN); exists {
		// 从Token中 -> 给UserId赋值
		req.UserId = value.(*token.Token).GetUserId()
	}

	_, err = h.service.PublishVideo(c, req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	ctx.JSON(http.StatusOK, custom.NewWithCode(constant.OPERATE_OK))
	return nil
}

func (h *Handler) publishList(c context.Context, ctx *app.RequestContext) error {
	req := video.NewPublishListRequest()
	// 1、接收参数
	if err := ctx.BindAndValidate(req); err != nil {
		h.log.Error(err)
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 业务请求
	videos, err := h.service.PublishList(c, req)
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

func (h *Handler) feed(c context.Context, ctx *app.RequestContext) error {
	req := video.NewFeedVideosRequest()
	// 1、接收参数
	if err := ctx.BindAndValidate(req); err != nil {
		h.log.Error(err)
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 业务请求
	videos, err := h.service.FeedVideos(c, req)
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
