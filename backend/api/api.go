package api

import (
	"fmt"
	"net/http"

	"gitlab.liu.se/adaab301/tddd27_2022_project/backend/api/middleware"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
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
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "GET", "POST"},
		AllowHeaders:     []string{"Content-Type", "Origin", "Authorization"},
		AllowCredentials: true,
	}))

	authorized := r.Group(ApiPath + "auth/")
	authorized.Use(gin.WrapH(middleware.EnsureValidToken()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		if err := middleware.HandleChunkUpload(r, token); err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err.Error()), 400)
		}
		w.Header().Add("Uid", token.RegisteredClaims.Subject)
	}))))
	{
		authorized.POST("videos/combine/", combineChunks)
		authorized.GET("videos/chunks/", chunkConstants)
		authorized.POST("videos/", uploadVideoChunk)
		authorized.POST("videos/save", saveVideo)
		authorized.POST("videos/comments/", addComment)
		authorized.GET("me/", getMe)
	}

	r.GET(ApiPath+"preview/", videoPreivew)
	r.GET(ApiPath+"video/", getVideo)
	r.GET(ApiPath+"videos/comments/", getComments)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func internalError(c *gin.Context, err error) {
	fmt.Printf(err.Error())
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}
