package api

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"gitlab.liu.se/adaab301/tddd27_2022_project/chunk"
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
	if err := chunk.CombineChunks(chunkName); err != nil {
		internalError(c, err)
	}
}
