package api

import (
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.liu.se/adaab301/tddd27_2022_project/transcode/transcoder/hls"
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

	file := c.Request.MultipartForm.File["videoFile"][0]
	dir := "tmp/" + file.Filename
	err = os.MkdirAll(dir, os.ModeTemporary)
	if err != nil {
		internalError(c, err)
		return
	}
	defer cleanup(dir)

	fileName := "tmp/" + file.Filename + "/" + file.Filename
	err = c.SaveUploadedFile(file, fileName)
	if err != nil {
		internalError(c, err)
		return
	}

	// Run FFMPEG
	err = hls.TranscodeToHLS(fileName, dir)
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
