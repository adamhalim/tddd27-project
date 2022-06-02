package chunk

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"time"

	"gitlab.liu.se/adaab301/tddd27_2022_project/lib/fileutil"
	"gitlab.liu.se/adaab301/tddd27_2022_project/lib/postgres"
)

const (
	transcoderUrl = "http://localhost:8081/api/transcode"
	saveUrl       = "http://localhost:8081/api/save"
	previewUrl    = "http://localhost:8081/video"
)

func ForwardVideoToTranscoder(chunkName string, fileName string, originalFileName string, uid string) error {
	// TODO: Might be worth to do this with gin instead of standard library
	// Uses multipart to upload file as chunks to transcoder
	client := &http.Client{}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("videoFile", fileName)
	if err != nil {
		return err
	}

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	_, err = io.Copy(fw, file)
	if err != nil {
		return err
	}

	writer.Close()
	req, err := http.NewRequest("POST", fmt.Sprintf("%s?chunkName=%s&originalFileName=%s&uid=%s", transcoderUrl, chunkName, url.QueryEscape(originalFileName), uid), bytes.NewReader(body.Bytes()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with code %d", res.StatusCode)
	}
	return nil
}

func SaveVideo(chunkName string, startTime float64, endTime float64, videoTitle string) error {
	client := &http.Client{}
	session, err := getSession(chunkName)
	if err != nil {
		return err
	}
	defer session.RemoveSession()
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?chunkName=%s&videoTitle=%s&start=%f&end=%f", saveUrl, chunkName, url.QueryEscape(videoTitle), startTime, endTime), nil)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with code %d", res.StatusCode)
	}
	if err := postgres.AddVideo(postgres.Video{
		Chunkname:  chunkName,
		LastViewed: time.Now().Unix(),
		Uid:        session.uid,
		ViewCount:  0,
	}); err != nil {
		return err
	}
	return nil
}

func GetVideoPreview(chunkName string) (string, error) {
	session, err := getSession(chunkName)
	if err != nil {
		return "", err
	}
	//TODO: implement actual reverse proxy.
	videoUrl := fmt.Sprintf("%s/%s/transcoded/%s.mp4", previewUrl, chunkName, fileutil.RemoveFileExtension(session.originalFileName))
	return videoUrl, nil

}
