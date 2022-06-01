package h264

import (
	"fmt"
	"os"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"gitlab.liu.se/adaab301/tddd27_2022_project/lib/fileutil"
	"gitlab.liu.se/adaab301/tddd27_2022_project/lib/objectstore"
	"gitlab.liu.se/adaab301/tddd27_2022_project/lib/videostats"
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

func CutVideo(fileName string, originalFileName string, videoTitle string, start float64, end float64, dir string, uid string, chunkName string) error {
	if err := validateStartAndEndTime(start, end); err != nil {
		return err
	}
	outputFileName := fmt.Sprintf("%s/%s.mp4", dir, fileutil.RemoveFileExtension(videoTitle))
	if start != 0 || end != 1 {
		frameCount, err := videostats.VideoFrameCount(fileName)
		if err != nil {
			return err
		}
		frameRate, err := videostats.VideoFramerate(fileName)
		if err != nil {
			return err
		}

		startFrame := start * float64(frameCount)
		endFrame := end * float64(frameCount)

		startTime := float64(startFrame) / frameRate
		endTime := float64(endFrame) / frameRate

		ffmpeg_go.Input(fileName).
			Output(outputFileName,
				ffmpeg_go.KwArgs{
					"c:v": "libx264",
					"ss":  startTime,
					"to":  endTime,
					"c":   "copy",
				}).
			OverWriteOutput().
			ErrorToStdOut().
			Run()
	} else {
		outputFileName = fileName
	}

	if err := objectstore.FiletoDB(fmt.Sprintf("%s/video.mp4", chunkName), outputFileName, originalFileName, uid, videoTitle); err != nil {
		return err
	}
	return nil
}

func validateStartAndEndTime(startTime float64, endTime float64) error {
	if startTime < 0 {
		return fmt.Errorf("error: invalid startTime of %.2f, must be >0", startTime)
	}
	if endTime > 1 {
		return fmt.Errorf("error: invalid endTime of %.2f, must be <= 1", endTime)
	}

	if startTime >= endTime {
		return fmt.Errorf("error: startTime (%.2f) must be less than endTime (%.2f)", startTime, endTime)
	}
	return nil
}
