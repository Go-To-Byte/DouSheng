// Author: BeYoung
// Date: 2023/2/19 23:03
// Software: GoLand

package milddles

import (
	"github.com/Go-To-Byte/DouSheng/network/models"
	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func Cos() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("data")
		if err != nil {
			zap.S().Debugf("failed to get form file: %v", err)
			c.JSON(http.StatusBadRequest, models.PublishResponse{
				StatusCode: 1,
				StatusMsg:  "failed",
			})
			c.Abort()
			return
		}

		u, _ := url.Parse(models.Config.CosConfig.Url)
		b := &cos.BaseURL{BucketURL: u}
		con := cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  os.Getenv("SECRETID"),
				SecretKey: os.Getenv("SECRETKEY"),
			},
		})

		VideoId := models.Node.Generate().Int64()
		c.Set("video_id", VideoId)

		// save file, because upload needed is a file, not an io.Reader
		dst := "~/log/api/video/" + strconv.FormatInt(VideoId, 10) + ".mp4"
		// 上传文件至指定的完整文件路径
		err = c.SaveUploadedFile(file, dst)
		defer func() {
			_ = os.Remove(dst)
		}()
		if err != nil {
			zap.S().Debugf("failed to save file: %v", err)
			c.JSON(http.StatusBadRequest, models.PublishResponse{
				StatusCode: 1,
				StatusMsg:  "failed",
			})
			c.Abort()
			return
		}

		// using upload file, because cos.upload could return video details
		name := "video/" + strconv.FormatInt(VideoId, 10) + ".mp4" // path = video/id.mp4
		details, _, err := con.Object.Upload(c, name, dst, nil)
		if err != nil {
			zap.S().Debugf("failed to upload video: %v", err)
			c.JSON(http.StatusBadRequest, models.PublishResponse{
				StatusCode: 1,
				StatusMsg:  "failed",
			})
			c.Abort()
			zap.S().Errorf("failed upload video: %v", err)
			return
		}

		// the url should get, because using the sample put
		ourl := con.Object.GetObjectURL(name)
		c.Set("video_url", details.Location)

		// get snapshot
		opt := cos.GetSnapshotOptions{Time: 1}
		snapshot, err := con.CI.GetSnapshot(c, name, &opt)
		if err != nil {
			c.Abort()
			zap.S().Debugf("failed to get snapshot: %v", err)
			return
		}

		// using put snapshot, because put is support io.Reader
		name = "images/" + strconv.FormatInt(VideoId, 10) + ".jpg"
		_, err = con.Object.Put(c, name, snapshot.Body, nil)
		if err != nil {
			zap.S().Debugf("failed to put snapshot: %v", err)
			c.JSON(http.StatusBadRequest, models.PublishResponse{
				StatusCode: 1,
				StatusMsg:  "failed",
			})
			c.Abort()
			zap.S().Errorf("failed upload snapshot: %v", err)
			return
		}
		ourl = con.Object.GetObjectURL(name)
		c.Set("cover_url", ourl)
	}
}
