import { Link } from 'react-router-dom'
import './style.css'


const Header = () => {
    return (
        <header>
            <div className="header-left">
            </div>
            <div className="header-mid">
            <Link to="/"><h1>vidds™</h1></Link>
            </div>
            <div className="header-right">
            </div>
        </header>
    )
}
export default Header