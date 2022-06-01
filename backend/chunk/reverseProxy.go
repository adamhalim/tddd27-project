package chunk

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

const (
	transcoderUrl = "http://localhost:8081/api/transcode"
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
