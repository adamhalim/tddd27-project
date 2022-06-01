package videostats

import (
	"encoding/json"
	"strconv"
	"strings"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

type probeData struct {
	Format probeDuration `json:"format"`
}

type probeDuration struct {
	Duration string `json:"duration"`
}

type probeStreams struct {
	Streams []interface{} `json:"streams"`
}

// Returns video duration in seconds
func VideoDuration(fileName string) (float64, error) {
	stats, err := ffmpeg_go.Probe(fileName, ffmpeg_go.KwArgs{})
	if err != nil {
		return 0, err
	}
	pd := probeData{}
	err = json.Unmarshal([]byte(stats), &pd)
	if err != nil {
		return 0, err
	}

	f, err := strconv.ParseFloat(pd.Format.Duration, 64)
	if err != nil {
		return 0, err
	}

	return f, nil
}

// Returns amount of frames in video
func VideoFrameCount(fileName string) (int64, error) {
	stats, err := ffmpeg_go.Probe(fileName, ffmpeg_go.KwArgs{})
	if err != nil {
		return 0, err
	}
	pi := probeStreams{}
	err = json.Unmarshal([]byte(stats), &pi)
	if err != nil {
		return 0, err
	}

	data := pi.Streams[0].(map[string]interface{})
	frameCount, err := strconv.ParseInt(data["nb_frames"].(string), 10, 0)
	if err != nil {
		return 0, err
	}

	return frameCount, nil
}

func VideoFramerate(fileName string) (float64, error) {
	ffmpeg_go.Input(fileName)
	_, err := VideoDuration(fileName)
	if err != nil {
		return 0, err
	}

	stats, err := ffmpeg_go.Probe(fileName, ffmpeg_go.KwArgs{})
	if err != nil {
		return 0, err
	}

	type probeFrameCount struct {
		Streams []struct {
			FrameCount int `json:"duration_ts"`
		} `json:"streams"`
	}
	ps := probeFrameCount{}
	err = json.Unmarshal([]byte(stats), &ps)
	if err != nil {
		return 0, err
	}

	type probeFramerate struct {
		Framerate []struct {
			Rate string `json:"r_frame_rate"`
		} `json:"streams"`
	}
	pf := probeFramerate{}
	err = json.Unmarshal([]byte(stats), &pf)
	if err != nil {
		return 0, err
	}

	splitFrameRate := strings.Split(pf.Framerate[0].Rate, "/")
	frameRate, err := strconv.Atoi(splitFrameRate[0])
	if err != nil {
		return 0, err
	}
	timeBase, err := strconv.Atoi(splitFrameRate[1])
	if err != nil {
		return 0, err
	}

	fps := float64(frameRate) / float64(timeBase)
	return fps, nil
}
