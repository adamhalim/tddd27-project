import './style.css'
type ProgressBarType = {
    progress: number
}
const ProgressBar = ({ progress }: ProgressBarType) => {
    const wrapperStyle = {
        height: '100%',
        width: `${progress}%`,
        backgroundColor: '#558B2F',
        textAlign: 'right' as 'right',
        borderRadius: 10
    }

    return (
        <div className="progress-bar-container">
            <div style={wrapperStyle}>
                <span>{`${progress.toFixed(0)}%`}</span>
            </div>
        </div>
    )
}



export default ProgressBar