package hls

import (
	"os"
	"sync"

	h "github.com/rendyfebry/go-hls-transcoder"
)

const (
	ffmpegPath = "ffmpeg"
)

func TranscodeToHLS(fileName string, dir string) error {

	targetPath := dir + "/hls"
	// TODO: Check source video res first and ignore resolutions greater than source
	resOptions := []string{"480p", "720p", "1080p"}

	err := os.MkdirAll(targetPath, os.ModeDir)
	if err != nil {
		return err
	}

	variants, err := h.GenerateHLSVariant(resOptions, "")
	if err != nil {
		return err
	}
	h.GeneratePlaylist(variants, targetPath, "")

	wg := sync.WaitGroup{}
	wg.Add(1)
	for _, res := range resOptions {
		defer wg.Done()
		err = h.GenerateHLS(ffmpegPath, fileName, targetPath, res)
		if err != nil {
			return err
		}
	}

	wg.Wait()

	return nil
}
