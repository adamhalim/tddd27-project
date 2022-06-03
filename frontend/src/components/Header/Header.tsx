import { useAuth0 } from '@auth0/auth0-react';
import { Link } from 'react-router-dom'
import AuthButtons from './AuthButtons/'
import './style.css'


const Header = () => {
    const { isAuthenticated } = useAuth0();

    return (
        <header>
            <div className="header-left">
            </div>
            <div className="header-mid">
            <Link to="/"><h1 className='text-4xl' >vidds™</h1></Link>
            </div>
            <div className="header-right">
                <AuthButtons 
                    isAuthenticated={isAuthenticated}
                />
            </div>
        </header>
    )
}
export default Header