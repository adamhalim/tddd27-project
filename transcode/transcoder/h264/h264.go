package h264

import (
	"fmt"
	"os"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"gitlab.liu.se/adaab301/tddd27_2022_project/lib/fileutil"
	"gitlab.liu.se/adaab301/tddd27_2022_project/lib/objectstore"
)

func TranscodeToh264(fileName string, originalFileName string, dir string, uid string) error {
	os.MkdirAll(fmt.Sprintf("%s/transcoded/", dir), os.ModeDir)
	outputFileName := fmt.Sprintf("%s/transcoded/%s.mp4", dir, fileutil.RemoveFileExtension(originalFileName))
	ffmpeg_go.Input(fileName).
		Output(outputFileName, ffmpeg_go.KwArgs{"c:v": "libx264"}).
		OverWriteOutput().
		ErrorToStdOut().
		Run()

	return nil
}
	return nil
}
