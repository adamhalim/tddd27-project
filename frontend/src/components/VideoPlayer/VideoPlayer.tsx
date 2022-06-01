import './style.css'

type VideoPlayerType = {
    videoSrc: string
}

const VideoPlayer = ({ videoSrc }: VideoPlayerType) => {
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
        </div>
    )
}

export default VideoPlayer