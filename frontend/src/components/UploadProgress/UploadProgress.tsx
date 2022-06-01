import { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'
import { fetchPreviewUrl } from '../../lib/fetchVideoURL'
import VideoTrimmer from '../VideoTrimmer'
import ProgressBar from './ProgressBar'
import './style.css'

type UploadProgressType = {
    progress: number,
    fileName: string,
    statusText: string,
    loading: boolean,
    errorOccured: boolean,
    chunkName: string,
    accessToken: string,
    videoSaved: boolean,
    saveVideo: (chunkName: string, start: number, end: number, title: string, accessToken: string, callback: VoidFunction) => Promise<boolean>
}

const UploadProgress = ({ progress, fileName, statusText, loading, errorOccured, chunkName, accessToken, videoSaved, saveVideo }: UploadProgressType) => {
    const [videoSrc, setVideoSrc] = useState("")

    const videoLink = `/video/${chunkName}`

    useEffect(() => {
        if (chunkName) {
            fetchPreviewUrl(chunkName).then((url) => {
                setVideoSrc(url as string)
            })
        } else {
            setVideoSrc("")
        }
    }, [chunkName])

    return (
        <div className='upload-progress-container'>
            <p>filename: {fileName}</p>
            <p
                className={`${loading ? "loading" : ""} ${errorOccured ? "error" : ""}`}
                style={{ fontFamily: 'monospace', display: 'inline-block' }}
            >
                {statusText}
            </p>
            {errorOccured ? <></> : <ProgressBar progress={progress} />}
            {videoSrc && <VideoTrimmer
                chunkName={chunkName}
                videoSrc={videoSrc}
                fileName={fileName}
                accessToken={accessToken}
                saveVideo={saveVideo}
            />}

            {videoSaved && <p>Link to video: <Link to={`/video/${chunkName}`}>{window.location.origin}{`/video/${chunkName}`}</Link></p>}

        </div>
    )
}

export default UploadProgress