package api

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.liu.se/adaab301/tddd27_2022_project/backend/chunk"
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
	fileName, directory, originalFileName, err := chunk.CombineChunks(chunkName)
	err = chunk.ForwardVideoToTranscoder(fileName, originalFileName)
	if err != nil {
		internalError(c, err)
	}
	os.RemoveAll(directory)
}

func chunkConstants(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"maxFileSize": chunk.MaxFileSize,
		"chunkSize":   chunk.ChunkSize,
	})
}
