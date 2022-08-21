import { video } from '../../pages/Profile/Profile';
import './style.css'
import VideoCard from './VideoCard';

type VideoGridType = {
    videos: video[],
    deleteVideo: (index:number) => void
    deletable: boolean,
}

const VideoGrid = ({ videos, deleteVideo, deletable }: VideoGridType) => {
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
                            deletable={deletable}
                        />
                    )
                }
            </div>
        </div>
    )
}

export default VideoGrid