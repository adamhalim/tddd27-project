import axios from 'axios';
import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom'
import VideoPlayer from '../../components/VideoPlayer';
import { fetchVideoURL } from '../../lib/fetchVideoURL';
import './style.css'

const instance = axios.create({
    baseURL: 'http://localhost:8080/api/video/',
});

interface videoStats {
    url: string
}

const VideoPage = () => {
    const { id } = useParams();
    const [loading, setLoading] = useState(true);
    const [videoURL, setVideoURL] = useState("");
    const [videoTitle, setVideoTitle] = useState("");
    const [viewCount, setViewCount] = useState(0);

    useEffect(() => {
        update()
    }, [])

    const update = async () => {
        if (id) {
            const data = await fetchVideoURL(id as string)
            if (data) {
                const { url, viewcount, videotitle } = data
                setVideoURL(url)
                setVideoTitle(videotitle)
                setViewCount(viewcount)
                setLoading(false)
            } else {
                // TODO: error handling
            }
        }
    }


    return (
        <div className='video-page-container'>
            {
                !loading &&
                <div className='video-page-player-wrapper'>
                    <div className='video-player-container'>
                        <VideoPlayer
                            videoSrc={videoURL}
                        />
                        <div className='video-player-stats'>
                            <div className='video-player-wrapper'>
                                <span className='video-player-title'> title: {videoTitle}</span>
                                <span className='video-player-viewcount'>viewcount: {viewCount} </span>
                            </div>
                        </div>
                    </div>


                </div>
            }

        </div>

    )
}

export default VideoPage