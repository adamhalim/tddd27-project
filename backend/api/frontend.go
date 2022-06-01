package api

import (
	"io/ioutil"
	"net/http"
	"os"
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
	}); err != nil {
		internalError(c, err)
		return
	}
	os.RemoveAll(directory)
}

func chunkConstants(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"maxFileSize": chunk.MaxFileSize,
		"chunkSize":   chunk.ChunkSize,
	})
}
