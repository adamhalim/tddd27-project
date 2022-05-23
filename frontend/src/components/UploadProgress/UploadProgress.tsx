import { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'
import { fetchVideoURL } from '../../lib/fetchVideoURL'
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
}

const UploadProgress = ({ progress, fileName, statusText, loading, errorOccured, chunkName }: UploadProgressType) => {
    const [videoSrc, setVideoSrc] = useState("")

    const videoLink = `/video/${chunkName}`

    useEffect(() => {
        if (chunkName) {
            fetchVideoURL(chunkName).then((url) => {
                setVideoSrc(url as string)
            })
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
            {chunkName && <p>Link to video: <Link to={videoLink}>{window.location.origin}{videoLink}</Link></p>}
            {chunkName && <VideoTrimmer videoSrc={videoSrc} fileName={fileName} />}
        </div>
    )
}

export default UploadProgress