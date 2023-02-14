// @Author: Ciusyan 2023/2/7
package api

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception/custom"
	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
	"github.com/Go-To-Byte/DouSheng/video_service/common/utils"
)

func (h *Handler) publishAction(ctx *gin.Context) error {

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
	if err = ctx.Bind(req); err != nil {
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}
	req.CoverUrl = uploaded.CoverRelativeURI
	req.PlayUrl = uploaded.RelativeURI
	_, err = h.service.PublishVideo(ctx.Request.Context(), req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	ctx.JSON(http.StatusOK, custom.NewWithCode(constant.OPERATE_OK))
	return nil
}

func (h *Handler) publishList(ctx *gin.Context) error {
	return nil
}

func (h *Handler) feed(ctx *gin.Context) error {
	return nil
}
