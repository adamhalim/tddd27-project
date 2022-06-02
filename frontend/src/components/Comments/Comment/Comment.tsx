import { timeAgo } from '../../../lib/getFormattedDate'
import './style.css'

type CommentType = {
    comment: string,
    author: string,
    date: Date,
}

const Comment = ({ comment, author, date }: CommentType) => {
    return (
        <div className='comment-container'>
            <div className='comment-author'>
                <span style={{ fontWeight: 'bold' }}>{author}</span>
                <span style={{ color: 'grey' }}> {timeAgo(date)}</span>
            </div>
            <div className='comment-content'>
                <span>{comment}</span>
            </div>
        </div>
    )
}
export default Comment