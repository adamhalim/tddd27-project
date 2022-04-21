package h264

import (
	"fmt"
	"strings"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"gitlab.liu.se/adaab301/tddd27_2022_project/transcode/db"
)

func TranscodeToh264(fileName string, originalFileName string, dir string, uid string) error {
	outputFileName := fmt.Sprintf("%s/%s.mp4", dir, removeFileExtension(originalFileName))
	ffmpeg_go.Input(fileName).
		Output(outputFileName, ffmpeg_go.KwArgs{"c:v": "libx264"}).
		OverWriteOutput().
		ErrorToStdOut().
		Run()

	db.AddFile(removeFileNameFromDirectory(dir)+"/"+removeFileNameFromDirectory(dir)+".mp4", outputFileName, originalFileName, uid)
	return nil
}

func removeFileExtension(fileName string) string {

	return strings.SplitN(fileName, ".", 2)[0]

}

func removeFileNameFromDirectory(dir string) string {
	return strings.SplitN(dir, "_", 2)[0][4:]
}
