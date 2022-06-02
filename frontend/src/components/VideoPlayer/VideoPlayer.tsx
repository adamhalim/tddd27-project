import './style.css'

type VideoPlayerType = {
    videoSrc: string,
    videoTitle: string,
    viewCount: number,
}

const VideoPlayer = ({ videoSrc, videoTitle, viewCount }: VideoPlayerType) => {
    return (
        <div className='video-player-container'>
            <video
                src={videoSrc}
                controls
                autoPlay
                muted
                loop
            >
            </video>
            <div className='video-player-stats'>
                <div className='video-player-wrapper'>
                    <span className='video-player-title'> title: {videoTitle}</span>
                    <span className='video-player-viewcount'>viewcount: {viewCount} </span>
                </div>
            </div>
        </div>
    )
}

export default VideoPlayer