import './style.css'

type VideoPlayerType = {
    videoSrc: string,
}

const VideoPlayer = ({ videoSrc }: VideoPlayerType) => {
    return (
        <>
            <video
                src={videoSrc}
                controls
                autoPlay
                muted
                loop
            >
            </video>
        </>
    )
}

export default VideoPlayer