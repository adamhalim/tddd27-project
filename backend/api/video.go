package api

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	endpoint        string
	accessKeyID     string
	secretAccessKey string
)

const (
	bucketName = "videos"
)

func init() {
	endpoint = os.Getenv("DB_ENDPOINT")
	if endpoint == "" {
		log.Fatal("no DB_ENDPOINT in .env")
	}

	accessKeyID = os.Getenv("ACCESS_KEY")
	if endpoint == "" {
		log.Fatal("no ACCESS_KEY in .env")
	}

	secretAccessKey = os.Getenv("SECRET_ACCESS_KEY")
	if endpoint == "" {
		log.Fatal("no SECRET_ACCESS_KEY in .env")
	}
}

func getMinioClient() *minio.Client {
	useSSL := false
	// Initialize minio client object.
	mc, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatal(err)
	}
	return mc
}

func GetVideo(c *gin.Context) {
	reqParams := make(url.Values)
	fileName := c.Query("fileName")
	reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	u, err := getMinioClient().PresignedGetObject(context.Background(), bucketName, fileName, time.Hour*1, reqParams)
	if err != nil {
		internalError(c, err)
	}
	c.JSON(200, gin.H{
		"videoUrl": u,
	})
}
