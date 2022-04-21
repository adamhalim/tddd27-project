import ProgressBar from './ProgressBar'
import './style.css'

type UploadProgressType = {
    progress: number,
    fileName: string,
    statusText: string,
    loading: boolean,
    errorOccured: boolean,
}

const UploadProgress = ({ progress, fileName, statusText, loading, errorOccured }: UploadProgressType) => {
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
        </div>
    )
}

export default UploadProgress