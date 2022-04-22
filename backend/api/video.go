package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.liu.se/adaab301/tddd27_2022_project/lib/objectstore"
)

var (
	endpoint        string
	accessKeyID     string
	secretAccessKey string
)

const (
	bucketName = "videos"
)

func getVideo(c *gin.Context) {
	chunkName := c.Query("chunkName")
	url, err := objectstore.GetVideoURL(chunkName)
	if err != nil {
		internalError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"url": url.String(),
	})
}
