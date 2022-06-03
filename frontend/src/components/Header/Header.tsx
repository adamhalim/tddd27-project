import { useAuth0 } from '@auth0/auth0-react';
import { Link } from 'react-router-dom'
import Profile from '../../pages/Profile';
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
                {
                    isAuthenticated && <div className='btn-header'>
                        <Link to='profile'>
                            <button>Profile</button>
                        </Link>
                    </div>
                }
                <AuthButtons
                    isAuthenticated={isAuthenticated}
                />

            </div>
        </header>
    )
}
export default Header