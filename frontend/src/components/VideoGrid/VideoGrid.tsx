import { video } from '../../pages/Profile/Profile';
import './style.css'
import VideoCard from './VideoCard';

type VideoGridType = {
    videos: video[],
    deleteVideo: (index:number) => void
}

const VideoGrid = ({ videos, deleteVideo }: VideoGridType) => {
    return (
        <div className="video-grid-container" >
            <div className='row'>
                {
                    videos.map((video, index) =>
                        <VideoCard
                            chunkName={video.Chunkname}
                            title={video.Title}
                            viewCount={video.ViewCount}
                            thumbnail='https://picsum.photos/300/200'
                            index={index}
                            deleteVideo={deleteVideo}
                        />
                    )
                }
            </div>
        </div>
    )
}

export default VideoGrid