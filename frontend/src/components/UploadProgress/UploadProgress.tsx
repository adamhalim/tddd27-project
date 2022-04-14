import ProgressBar from './ProgressBar'
import './style.css'

type UploadProgressType = {
    progress: number,
    fileName: string,
}

const UploadProgress = ({ progress, fileName }: UploadProgressType) => {
    return (
        <div className='upload-progress-container'>
            <p>filename: {fileName}</p>
            <ProgressBar progress={progress} />
        </div>
    )
}

export default UploadProgress