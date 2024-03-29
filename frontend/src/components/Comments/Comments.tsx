import { useAuth0 } from '@auth0/auth0-react'
import axios from 'axios'
import { useEffect, useState } from 'react'
import Comment from './Comment/Comment'
import './style.css'

type CommentsType = {
    chunkName: string,
    loadAccessToken: () => Promise<String>
}

interface Comment {
    comment: string,
    author: string,
    date: Date,
}

const Comments = ({ chunkName, loadAccessToken }: CommentsType) => {
    const [accessToken, setAccessToken] = useState("");
    const [comments, setComments] = useState<Comment[]>([])
    const [commentsCount, setCommentsCount] = useState(0)
    const [newComment, setNewComment] = useState("")

    const { getAccessTokenSilently, getAccessTokenWithPopup, user } = useAuth0();

    useEffect(() => {
        loadAccessToken().then((res) => {
            if (typeof res === 'string') {
                setAccessToken(res)
            } else {
                // TODO: Handle error ??
            }
        });
        loadComments()
    }, [])


    const loadComments = async () => {
        const res = await axios.get('http://localhost:8080/api/videos/comments/', {
            params: {
                chunkName: chunkName,
            }
        })
        if (res.status === 200) {
            interface response {
                Comment: string,
                Username: string,
                Date: number
            }
            const data = res.data.data as response[]
            console.log(data)
            let comments: Comment[] = []
            data.forEach((c) => {
                const date = new Date(c.Date)
                comments.push({ comment: c.Comment, author: c.Username, date: date })
            })
            setComments(comments)
            setCommentsCount(comments.length)
        }
    }

    const changeNewComment = (event: any) => {
        setNewComment(event.target.value)
    }

    const addComment = async () => {
        if (!!accessToken) {
            await loadAccessToken()
        }
        const res = await axios.post('http://localhost:8080/api/auth/videos/comments/', {
            chunkname: chunkName,
            comment: newComment,
        }, {
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${accessToken}`
            },
            withCredentials: true,
        })
        if (res.status === 201 && user?.sub) {
            loadComments()
            setNewComment("")
        } else {
            // TODO: handle error :D
        }
    }

    return (
        <div className='comments-container'>
            <div className='comments-header' >
                <span>{commentsCount} comments</span>
            </div>
            <div className='comments-content'>
                {
                    comments.map(({ comment, author, date }, index) =>
                        <Comment
                            key={index}
                            comment={comment}
                            author={author}
                            date={date}
                        />
                    )
                }
                {
                    accessToken && <div className='comments-add'>
                        <input
                            placeholder='Add a comment...'
                            onChange={changeNewComment}
                            value={newComment}
                        >
                        </input>
                        <button
                            onClick={addComment}
                            disabled={!newComment}
                        >
                            Comment
                        </button>
                    </div>
                }

            </div>
        </div>
    )
}

export default Comments