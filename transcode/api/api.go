package api

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	ApiPath = "/api/"
)

func Start() {
	handleRequests()
}

func handleRequests() {
	r := gin.Default()

	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{"Content-Type", "Origin"},
	}))

	r.POST(ApiPath+"transcode/", postVideo)
	r.GET(ApiPath+"save/", saveVideo)

	r.Run(":8081")
}

func internalError(c *gin.Context, err error) {
	fmt.Printf(err.Error())
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}
