package api

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.liu.se/adaab301/tddd27_2022_project/backend/chunk"
	"gitlab.liu.se/adaab301/tddd27_2022_project/lib/objectstore"
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

func addComment(c *gin.Context) {
	type commentData struct {
		ChunkName string `json:"chunkname"`
		Comment   string `json:"comment"`
		AuthorUid string
	}
	var comment commentData
	if err := c.BindJSON(&comment); err != nil {
		internalError(c, err)
		return
	}

	uid := gin.ResponseWriter.Header(c.Writer)["Uid"][0]
	if uid == "" {
		internalError(c, errors.New("no uid provided"))
		return
	}
	comment.AuthorUid = uid

	err := postgres.AddComment(comment.ChunkName, comment.Comment, comment.AuthorUid)
	if err != nil {
		internalError(c, err)
		return
	}

	c.Status(http.StatusCreated)
}

func getComments(c *gin.Context) {
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

	comments, err := postgres.GetComments(chunkName)
	if err != nil {
		internalError(c, err)
		return
	}
	if chunkName == "" {
		internalError(c, errors.New("no chunkName provided"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": comments,
	})
}

func getMe(c *gin.Context) {
	uid := gin.ResponseWriter.Header(c.Writer)["Uid"][0]
	if uid == "" {
		internalError(c, errors.New("no uid provided"))
		return
	}

	user, err := postgres.FindUser(uid)
	if err != nil {
		internalError(c, err)
		return
	}

	videos, err := postgres.FindVideosFromUser(uid)
	if err != nil {
		internalError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"videos":   videos,
	})
}

func changeUsername(c *gin.Context) {
	uid := gin.ResponseWriter.Header(c.Writer)["Uid"][0]
	if uid == "" {
		internalError(c, errors.New("no uid provided"))
		return
	}

	_, err := postgres.FindUser(uid)
	if err != nil {
		internalError(c, err)
		return
	}

	queryParams, err := url.ParseQuery(c.Request.URL.RawQuery)
	if err != nil {
		internalError(c, err)
		return
	}
	username := queryParams["username"][0]
	if username == "" {
		internalError(c, errors.New("no username provided"))
		return
	}

	err = postgres.ChangeUsername(uid, username)
	if err != nil {
		internalError(c, err)
		return
	}

	c.Status(http.StatusOK)

}

func deleteVideo(c *gin.Context) {
	queryParams, err := url.ParseQuery(c.Request.URL.RawQuery)
	if err != nil {
		internalError(c, err)
		return
	}
	uid := gin.ResponseWriter.Header(c.Writer)["Uid"][0]
	if uid == "" {
		internalError(c, errors.New("no uid provided"))
		return
	}
	chunkName := queryParams["chunkName"][0]
	if chunkName == "" {
		internalError(c, errors.New("no chunkName provided"))
		return
	}

	err = postgres.DeleteVideo(uid, chunkName)
	if err != nil {
		internalError(c, err)
		return
	}

	err = objectstore.DeleteVideo(uid, chunkName)
	if err != nil {
		internalError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func likeVideo(c *gin.Context) {
	queryParams, err := url.ParseQuery(c.Request.URL.RawQuery)
	if err != nil {
		internalError(c, err)
		return
	}
	uid := gin.ResponseWriter.Header(c.Writer)["Uid"][0]
	if uid == "" {
		internalError(c, errors.New("no uid provided"))
		return
	}
	chunkName := queryParams["chunkName"][0]
	if chunkName == "" {
		internalError(c, errors.New("no chunkName provided"))
		return
	}

	err = postgres.LikeVideo(chunkName, uid)
	if err != nil {
		// If the video has already been liked
		if (err.(*pq.Error)).Code == "23505" {
			c.Status(http.StatusOK)
			return
		}
		internalError(c, err)
		return
	}

	c.Status(http.StatusOK)
}

		internalError(c, err)
		return
	}

	c.Status(http.StatusOK)
}
