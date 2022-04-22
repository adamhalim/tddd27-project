package api

import (
	"fmt"

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
	url, err := objectstore.GetVideo(chunkName)
	if err != nil {
		internalError(c, err)
	}

	fmt.Printf("%v", url)
}
