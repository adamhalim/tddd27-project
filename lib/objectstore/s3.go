package objectstore

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"gitlab.liu.se/adaab301/tddd27_2022_project/lib/fileutil"
)

var (
	endpoint        string
	accessKeyID     string
	secretAccessKey string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	endpoint = os.Getenv("DB_ENDPOINT")
	if endpoint == "" {
		log.Fatal("no DB_ENDPOINT in .env")
	}

	accessKeyID = os.Getenv("ACCESS_KEY")
	if accessKeyID == "" {
		log.Fatal("no ACCESS_KEY in .env")
	}

	secretAccessKey = os.Getenv("SECRET_ACCESS_KEY")
	if secretAccessKey == "" {
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
		log.Fatalln(err)
	}
	return mc
}

const (
	bucketName = "videos"
)

func FilesFromDirectory(originalFileName string, dir string, uid string) error {
	dirName := fileutil.RemoveFileNameFromDirectory(dir)
	// HLS files are stored in tmp/chunkname_originalFilename/hls
	filepath.Walk(dir+"/hls", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		// We add all files to the bucket at chunkName/file.ts
		FiletoDB(dirName+"/"+info.Name(), path, originalFileName, uid, originalFileName)
		return nil
	})
	return nil
}

func FiletoDB(fileName string, filePath string, originalFileName string, uid string, videoTitle string) error {
	if _, err := getMinioClient().FPutObject(context.Background(), bucketName, fileName, filePath, minio.PutObjectOptions{
		ContentType: "application/video",
	}); err != nil {
		return err
	}
	return nil
}

func GetVideoURL(chunkName string) (*url.URL, error) {
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=\"%s/video.mp4\"", chunkName))
	u, err := getMinioClient().PresignedGetObject(context.Background(), bucketName, chunkName+"/video.mp4", time.Hour*1, reqParams)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func DeleteVideo(uid string, chunkName string) error {

	err := getMinioClient().RemoveObject(context.Background(), bucketName, chunkName+"/video.mp4", minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
