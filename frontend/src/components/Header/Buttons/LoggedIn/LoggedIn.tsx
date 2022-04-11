import { Link } from "react-router-dom"

const LoggedIn = () => {
    return (
        <div>
            <Link to='/profile' role="button" className="btn-header">Profile</Link>
            <Link to='/' role="button" className="btn-header">Log out</Link>
        </div>
    )
}

export default LoggedIn