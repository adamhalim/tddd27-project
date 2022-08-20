import { useAuth0 } from '@auth0/auth0-react'
import axios from 'axios';
import { useEffect, useState } from 'react';
import VideoGrid from '../../components/VideoGrid';
import './style.css'

export interface video {
    Chunkname: string,
    ViewCount: number,
    Title: string,
}

const Profile = () => {
    const { isAuthenticated, loginWithRedirect, getAccessTokenSilently, getAccessTokenWithPopup } = useAuth0();
    const [username, setUsername] = useState('')
    const [videos, setVideos] = useState<video[]>([])
    const [accessToken, setAccessToken] = useState('')
    const [editingUsername, setEditingUsername] = useState(false)


    useEffect(() => {
        if (!isAuthenticated) {
            loginWithRedirect()
        }
        loadAccessToken()
    }, [])
    useEffect(() => {
        if (accessToken)
            getMe()
    }, [accessToken])

    const loadAccessToken = async () => {
        const token = await getAccessTokenSilently({ audience: 'http://localhost:3000/' })
            .then((res) => {
                return res;
            }).catch((err) => {
                console.log(err);
                // getAccessTokenSilently() with audience won't work on localhost,
                // but will work with a popup. Ghetto workaround, but it works for now..
                return getAccessTokenWithPopup({ audience: 'http://localhost:3000/' })
            })
        setAccessToken(token)
    }


    const getMe = async () => {
        const res = await axios.get('http://localhost:8080/api/auth/me/', {
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${accessToken}`
            },
            withCredentials: true,
        })

        interface response {
            username: string,
            videos: video[],
        }

        const data = res.data as response
        setUsername(data.username)
        setVideos(data.videos)
    }

    const deleteVideo = async (index: number) => {
        const chunkname = videos[index].Chunkname
        const res = await axios.delete('http://localhost:8080/api/auth/videos/', {
            params: {
                "chunkName": chunkname,
            },
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${accessToken}`
            },
            withCredentials: true,
        })
        if (res.status !== 204) {
            console.log('error:', res)
        } else {
            setVideos(videos.filter((_, i) => {
                return i !== index;
            }));
            // TODO: Alert that video has been deleted
        }
    }

    const handleEditUsername = async () => {
        if (editingUsername) {
            await saveUsername()
        }
        setEditingUsername(!editingUsername)
    }

    const changeUsername = async (event: any) => {
        setUsername(event.target.value)
    }

    const saveUsername = async () => {
        const res = await axios.post('http://localhost:8080/api/auth/username/', {}, {
            params: {
                username: username,
            },
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${accessToken}`
            },
            withCredentials: true,
        })
        if (res.status !== 200) {
            console.log(res.statusText)
            // TODO: Handle error :D
        }
    }

    if (!username && !accessToken) {
        return <div></div>
    }

    return (
        <div className='profile-container' >
            <div className='profile-card'>
                <div className='profile-username'>
                    {
                        editingUsername ? <div>
                            <span>username: </span>
                            <input
                                value={username}
                                onChange={changeUsername}
                            />
                        </div> :
                            <div>
                                <span>username: {username}</span>
                            </div>
                    }

                    <button
                        onClick={handleEditUsername}
                    >
                        &#128393;
                    </button>
                </div>
                {videos && <VideoGrid
                    videos={videos}
                    deleteVideo={deleteVideo}
                />}
            </div>
        </div>
    )
}

export default Profile