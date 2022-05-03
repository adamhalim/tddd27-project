import axios from 'axios';
import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom'
import VideoPlayer from '../../components/VideoPlayer';
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

    useEffect(() => {
        fetchVideoURL()
    }, [])

    const fetchVideoURL = async () => {
        const res = await instance.get('', {
            params: {
                chunkName: id,
            }
        })
        setLoading(false)
        if (res.status === 200) {
            const data: videoStats = res.data
            setVideoURL(data.url)
        } else {
            // TODO: Error handling here
        }
    }

    return (
        <div className='video-page-container'>
            <div className='video-page-player-wrapper'>
                {
                    !loading && <VideoPlayer videoSrc={videoURL} />
                }
            </div>

        </div>

    )
}

export default VideoPage