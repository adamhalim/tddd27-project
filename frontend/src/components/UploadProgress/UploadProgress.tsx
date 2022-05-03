import { Link } from 'react-router-dom'
import ProgressBar from './ProgressBar'
import './style.css'

type UploadProgressType = {
    progress: number,
    fileName: string,
    statusText: string,
    loading: boolean,
    errorOccured: boolean,
    videoURL: string,
}

const UploadProgress = ({ progress, fileName, statusText, loading, errorOccured, videoURL }: UploadProgressType) => {
    return (
        <div className='upload-progress-container'>
            <p>filename: {fileName}</p>
            <p 
                className={`${loading ? "loading" : ""} ${errorOccured ? "error" : ""}`}
                style={{fontFamily:'monospace', display:'inline-block'}}
            >
                {statusText}
            </p>
            { errorOccured ? <></> : <ProgressBar progress={progress} /> } 
            { videoURL && <p>Link to video: <Link to={videoURL}>{window.location.origin}{videoURL}</Link></p>}
        </div>
    )
}

export default UploadProgress