import ProgressBar from './ProgressBar'
import './style.css'

type UploadProgressType = {
    progress: number,
    fileName: string,
    statusText: string,
    loading: boolean
}

const UploadProgress = ({ progress, fileName, statusText, loading: loading }: UploadProgressType) => {
    return (
        <div className='upload-progress-container'>
            <p>filename: {fileName}</p>
            <p 
                className={`${loading ? "loading" : ""}`}
                style={{fontFamily:'monospace', display:'inline-block'}}
            >
                {statusText}
            </p>
            <ProgressBar progress={progress} />
        </div>
    )
}

export default UploadProgress