import { useAuth0 } from '@auth0/auth0-react';
import axios from 'axios';
import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom'
import Comments from '../../components/Comments';
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

    const { getAccessTokenSilently, getAccessTokenWithPopup } = useAuth0();

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

    const loadAccessToken = async (): Promise<String> => {
        const accessToken = await getAccessTokenSilently({ audience: 'http://localhost:3000/' })
            .then((res) => {
                return res;
            }).catch((err) => {
                console.log(err);
                // getAccessTokenSilently() with audience won't work on localhost,
                // but will work with a popup. Ghetto workaround, but it works for now..
                return getAccessTokenWithPopup({ audience: 'http://localhost:3000/' })
            })
        return accessToken;
    }

    const likeVideo = () => {
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
                        <div className='video-page-stats'>
                            <div className='video-page-title'> title: {videoTitle}</div>
                            <div className='video-page-viewcount'>viewcount: {viewCount} likes: 0 <button>&hearts;</button> </div>
                        </div>
                        <div className='video-page-comments'>
                            <Comments
                                chunkName={id as string}
                                loadAccessToken={loadAccessToken}
                            />
                        </div>
                    </div>

                </div>
            }

        </div>

    )
}

export default VideoPage