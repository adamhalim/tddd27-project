import './style.css'
type VideoCardType = {
    thumbnail: string,
    chunkName: string,
    viewCount: number,
    title: string,
    index: number,
    deleteVideo: (index:number) => void
}
const VideoCard = ({ thumbnail, chunkName, viewCount, title, index, deleteVideo }: VideoCardType) => {

    return (
        <div className='video-card-elem col-sm-8 col-md-4 col-lg-2'>
            <a href={`video/${chunkName}`} target='_blank'>
                <img className='video-card-thumbnail'
                    src={thumbnail}
                />
            </a>

            <div className='video-card-block'>
                <p><span style={{ fontWeight: 'bold' }}>Title: </span>{title}</p>
                <div className='btn-video-card'>
                    <a href={`video/${chunkName}`} target='_blank'>View</a> | <a href='#' onClick={() => deleteVideo(index)}>Delete</a> | <span>{viewCount} views</span>
                </div>
            </div>
        </div>
    )
}

VideoCard.defaultProps = {
    thumbnail: "https://picsum.photos/300/200"
}

export default VideoCard