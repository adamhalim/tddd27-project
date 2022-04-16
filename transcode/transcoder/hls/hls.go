package hls

import (
	"os"

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

	for _, res := range resOptions {
		err = h.GenerateHLS(ffmpegPath, fileName, targetPath, res, true)
		if err != nil {
			return err
		}
	}
	return nil
}
