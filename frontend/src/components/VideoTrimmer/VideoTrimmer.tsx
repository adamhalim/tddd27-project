import './style.css'
import Nouislider from 'nouislider-react'
import "nouislider/distribute/nouislider.css";
import { useEffect, useState } from 'react';
import LoadingSpinner from '../LoadingSpinner';
type VideoTrimmer = {
    videoSrc: string,
    fileName: string,
    chunkName: string,
    accessToken: string,
    saveVideo: (chunkName: string, start: number, end: number, title: string, accesstoken: string, callback: VoidFunction) => Promise<boolean>
}
const VideoTrimmer = ({ chunkName, videoSrc, fileName, accessToken, saveVideo }: VideoTrimmer) => {
    let videoElement: HTMLVideoElement;

    const [previousStartTime, setPreviousStartTime] = useState(0)
    const [previousEndTime, setPreviousEndTime] = useState(1)
    const [startTime, setStartTime] = useState(0)
    const [endTime, setEndTime] = useState(1)
    const [videoDuration, setVideoDuration] = useState(0)
    const [videoTitle, setVideoTitle] = useState(fileName)
    const [videoTitleIsEdited, setVideoTitleIsEdited] = useState(false)
    const [waitingForSave, setWaitingForSave] = useState(false)
    const [saved, setSaved] = useState(false)

    useEffect(() => {
        if (videoElement?.duration) {
            setVideoDuration(videoElement.duration)
            if (startTime !== previousStartTime) {
                videoElement.currentTime = startTime * videoElement.duration
            } else if (endTime !== previousEndTime) {
                videoElement.currentTime = endTime * videoElement.duration
            }
        }
        setPreviousStartTime(startTime)
        setPreviousEndTime(endTime)
    }, [startTime, endTime])

    const handleSliderUpdate = (event: any) => {
        const start = parseFloat(event[0])
        const end = parseFloat(event[1])
        setStartTime(start)
        setEndTime(end)
    }

    const handleVideoMounted = (element: HTMLVideoElement) => {
        if (element !== null) {
            videoElement = element
        }
    }

    const handleTitleChange = (event: any) => {
        setVideoTitle(event.target.value)
    }

    const handleSave = () => {
        // TODO: send start, end & title to transcoder
        setWaitingForSave(true)
        saveVideo(chunkName, startTime, endTime, videoTitle, accessToken, videoIsSaved)
    }

    const videoIsSaved = () => {
        setWaitingForSave(false)
        setSaved(true)
    }

    return (
        <div className='video-trimmer-container'>
            <video
                src={videoSrc}
                controls
                autoPlay
                muted
                loop
                ref={handleVideoMounted}
            >
            </video>
            <Nouislider
                start={[0, 1]}
                range={{
                    min: 0,
                    max: 1
                }}
                margin={0.05}
                onUpdate={handleSliderUpdate}
                connect={true}
            />
            <div className='video-trimmer-wrapper'>
                <span>start: {(startTime * videoDuration).toFixed(1)}s</span>
            </div>
            <div className='video-trimmer-wrapper'>
                <span>end: {(endTime * videoDuration).toFixed(1)}s</span>
            </div>
            <div className='video-trimmer-title'>
                <div className='video-trimmer-wrapper'>
                    <span>Title: </span>
                </div>
                <div
                    onClick={() => setVideoTitleIsEdited(true)}
                >
                    {
                        videoTitleIsEdited ?
                            <div className='video-trimmer-wrapper'>
                                <input
                                    value={videoTitle}
                                    onChange={handleTitleChange}
                                    autoFocus

                                ></input>
                                <button className='video-trimmer-title-save'
                                    onClick={() => setTimeout(() => {
                                        setVideoTitleIsEdited(false)
                                    }, 0)}
                                /* No idea why I have to use setTimeout, it won't work by just setting state directly*/
                                >
                                    &#10003;
                                </button>
                            </div>
                            :
                            <div style={{ display: 'flex', overflowWrap: 'break-word' }}>
                                {videoTitle}
                                <div style={{ fontStyle: 'italic', color: 'grey' }}>
                                </div>
                            </div>
                    }
                </div>

            </div>
            <div className='video-trimmer-save' >
                <div>
                    {!saved &&
                        <button
                            onClick={handleSave}
                        >
                            Save
                        </button>
                    }
                </div>
                <div>
                    {
                        waitingForSave && <LoadingSpinner />
                    }
                </div>
            </div>
        </div>
    )
}

export default VideoTrimmer