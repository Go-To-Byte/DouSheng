// @Author: Ciusyan 2023/2/7
package api

import (
	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
	"github.com/Go-To-Byte/DouSheng/video_service/common/utils"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/Go-To-Byte/DouSheng/dou_common/constant"

	videoconst "github.com/Go-To-Byte/DouSheng/video_service/common/constant"
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

	ctx.JSON(http.StatusOK, uploaded)
}

func (h *Handler) publishList(ctx *gin.Context) {

}

func (h *Handler) feed(ctx *gin.Context) {

}
