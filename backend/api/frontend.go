package api

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.liu.se/adaab301/tddd27_2022_project/backend/chunk"
	"gitlab.liu.se/adaab301/tddd27_2022_project/lib/postgres"
)

func uploadVideoChunk(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	id := c.Query("id")
	fileName := c.Query("fileName")
	chunkName := c.Query("chunkName")

	if len(body) == 0 {
		// ERR
		return
	}

	if err := chunk.CreateChunk(body, id, fileName, chunkName); err != nil {
		internalError(c, err)
	}
}

func combineChunks(c *gin.Context) {
	chunkName := c.Query("chunkName")
	fileName, directory, originalFileName, uid, err := chunk.CombineChunks(chunkName)
	if err != nil {
		internalError(c, err)
		return
	}
	err = chunk.ForwardVideoToTranscoder(chunkName, fileName, originalFileName, uid)
	if err != nil {
		internalError(c, err)
		return
	}
	if err := postgres.AddVideo(postgres.Video{
		Chunkname:  chunkName,
		LastViewed: time.Now().Unix(),
		Uid:        uid,
		ViewCount:  0,
	}); err != nil {
		internalError(c, err)
		return
	}
	os.RemoveAll(directory)
}

func saveVideo(c *gin.Context) {
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
	videoTitle := queryParams["videoTitle"][0]
	if videoTitle == "" {
		internalError(c, errors.New("no videoTitle provided"))
		return
	}
	startString := queryParams["start"][0]
	if startString == "" {
		internalError(c, errors.New("no start provided"))
		return
	}
	endString := queryParams["end"][0]
	if endString == "" {
		internalError(c, errors.New("no end provided"))
		return
	}
	start, err := strconv.ParseFloat(startString, 64)
	if err != nil {
		internalError(c, err)
		return
	}
	end, err := strconv.ParseFloat(endString, 64)
	if err != nil {
		internalError(c, err)
		return
	}

	if err := chunk.SaveVideo(chunkName, start, end, videoTitle); err != nil {
		internalError(c, err)
		return
	}
}

func chunkConstants(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"maxFileSize": chunk.MaxFileSize,
		"chunkSize":   chunk.ChunkSize,
	})
}
