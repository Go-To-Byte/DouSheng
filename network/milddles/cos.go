// Author: BeYoung
// Date: 2023/2/19 23:03
// Software: GoLand

package milddles

import (
	"fmt"
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

		// using put file, because cos.upload not support put file by io.Reader
		read, err := file.Open()
		name := "video/" + strconv.FormatInt(VideoId, 10) + ".mp4" // path = video/id.mp4
		_, err = con.Object.Put(c, name, read, nil)
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
		c.Set("video_url", fmt.Sprint(ourl))

		// get snapshot
		opt := cos.GetSnapshotOptions{Time: 1}
		snapshot, err := con.CI.GetSnapshot(c, name, &opt)
		if err != nil {
			c.Abort()
			zap.S().Debugf("failed to get snapshot: %v", err)
			return
		}

		// using put snapshot, because put is support io.Reader
		name = "image/" + strconv.FormatInt(VideoId, 10) + ".jpg"
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
		c.Set("cover_url", fmt.Sprint(ourl))
	}
}
