package h264

import (
	"fmt"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"gitlab.liu.se/adaab301/tddd27_2022_project/transcode/fileutil"
	"gitlab.liu.se/adaab301/tddd27_2022_project/transcode/upload"
)

func TranscodeToh264(fileName string, originalFileName string, dir string, uid string) error {
	outputFileName := fmt.Sprintf("%s/%s.mp4", dir, fileutil.RemoveFileExtension(originalFileName))
	ffmpeg_go.Input(fileName).
		Output(outputFileName, ffmpeg_go.KwArgs{"c:v": "libx264"}).
		OverWriteOutput().
		ErrorToStdOut().
		Run()

	dirName := fileutil.RemoveFileNameFromDirectory(dir)
	upload.FiletoDB(dirName+"/"+"video"+".mp4", outputFileName, originalFileName, uid)
	return nil
}
