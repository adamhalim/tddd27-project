import { Link } from "react-router-dom"
import '../style.css'

const LoggedOut = () => {
    return (
        <div>
            <Link to='/login' role="button" className="btn-header" >Log in</Link>
            <Link to="/register" role="button" className="btn-header" >Register</Link>
        </div>
    )
}

export default LoggedOut