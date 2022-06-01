package api

import (
	"errors"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.liu.se/adaab301/tddd27_2022_project/transcode/transcoder/h264"
)

const (
	maxMemory = 32 << 20 // 32MB
)

type Form struct {
	Value map[string][]string
	File  map[string][]*multipart.FileHeader
}

func init() {
	os.RemoveAll("tmp/")
	err := os.MkdirAll("tmp/", os.ModeDir)
	if err != nil {
		log.Fatal(err)
	}
}

// Endpoint that downloads a video and transcodes it.
//
// Could be exported and wrapped in an AWS lambda.
func postVideo(c *gin.Context) {
	err := c.Request.ParseMultipartForm(maxMemory)
	if err != nil {
		internalError(c, err)
		return
	}
	queryParams, err := url.ParseQuery(c.Request.URL.RawQuery)
	if err != nil {
		internalError(c, err)
		return
	}
	originalFileName := queryParams["originalFileName"][0]
	if originalFileName == "" {
		internalError(c, errors.New("no originalFileName provided"))
		return
	}
	uid := queryParams["uid"][0]
	if uid == "" {
		internalError(c, errors.New("no uid provided"))
		return
	}

	file := c.Request.MultipartForm.File["videoFile"][0]
	dir := "tmp/" + chunkName
	err = os.MkdirAll(dir, os.ModeTemporary)
	if err != nil {
		internalError(c, err)
		return
	}

	fileName := dir + file.Filename
	err = c.SaveUploadedFile(file, fileName)
	if err != nil {
		internalError(c, err)
		return
	}

	// Run FFMPEG
	err = h264.TranscodeToh264(fileName, originalFileName, dir, uid)
	if err != nil {
		internalError(c, err)
		return
	}

	c.Status(http.StatusOK)

	// TODO: upload to database & do
}

func cleanup(directory string) {
	os.RemoveAll(directory)
}
