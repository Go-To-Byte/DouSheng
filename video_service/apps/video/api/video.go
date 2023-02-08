// @Author: Ciusyan 2023/2/7
package api

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"

	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
	videoconst "github.com/Go-To-Byte/DouSheng/video_service/common/constant"
	"github.com/Go-To-Byte/DouSheng/video_service/common/utils"
)

func (h *Handler) publishAction(ctx *gin.Context) {

	// 1、读取文件数据并且上传
	file, err := ctx.FormFile(videoconst.REQUEST_FILE)
	if err != nil {
		ctx.JSON(http.StatusNotFound, videoconst.BAD_NO_FILE)
		return
	}

	uploaded, err := utils.UploadFile(file)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, videoconst.BAD_UPLOAD_FILE)
		return
	}
	req := video.NewPublishVideoRequest()
	// 2、接收参数
	if err = ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, constant.ERROR_ARGS_VALIDATE)
		return
	}
	req.CoverUrl = uploaded.CoverRelativeURI
	req.PlayUrl = uploaded.RelativeURI

	_, err = h.service.PublishVideo(ctx.Request.Context(), req)
	if err != nil {
		e := err.(exception.CustomException)
		ctx.JSON(http.StatusBadRequest, e.GetCodeMsg())
	}

	ctx.JSON(http.StatusOK, constant.OPERATE_OK)
}

func (h *Handler) publishList(ctx *gin.Context) {

}

func (h *Handler) feed(ctx *gin.Context) {

}
