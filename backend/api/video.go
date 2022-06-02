package api

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
	"gitlab.liu.se/adaab301/tddd27_2022_project/backend/chunk"
	"gitlab.liu.se/adaab301/tddd27_2022_project/lib/objectstore"
	"gitlab.liu.se/adaab301/tddd27_2022_project/lib/postgres"
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

	err = postgres.IncrementViewCount(chunkName)
	if err != nil {
		internalError(c, err)
		return
	}

	video, err := postgres.FindVideo(chunkName)
	if err != nil {
		internalError(c, fmt.Errorf("no video found with chunkName %s", chunkName))
		return
	}

	c.JSON(200, gin.H{
		"url":        url.String(),
		"viewcount":  video.ViewCount,
		"videotitle": video.Title,
	})
}

func videoPreivew(c *gin.Context) {
	queryParams, err := url.ParseQuery(c.Request.URL.RawQuery)
	if err != nil {
		internalError(c, err)
		return
	}
	chunkName := queryParams["chunkName"][0]
	if chunkName == "" {
		internalError(c, errors.New("no chunkName provided"))
		return
	}
	url, err := chunk.GetVideoPreview(chunkName)
	if err != nil {
		internalError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"url": url,
	})
}
