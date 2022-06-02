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

    useEffect(() => {
        update()
    }, [])

    const update = async () => {
        if (id) {
            const url = await fetchVideoURL(id as string)
            if (typeof url === 'string') {
                setVideoURL(url)
                setLoading(false)
            } else {
                // TODO: error handling
            }
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