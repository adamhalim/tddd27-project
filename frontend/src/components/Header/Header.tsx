import { Link } from 'react-router-dom'
import LoggedIn from './Buttons/LoggedIn'
import LoggedOut from './Buttons/LoggedOut'
import './style.css'

type HeaderType = {
    signedIn: boolean,
    signIn: VoidFunction
}

const Header = ({ signedIn, signIn }: HeaderType) => {
    return (
        <header>
            <div className="header-left">
            </div>
            <div className="header-mid">
            <Link to="/"><h1 className='text-4xl' >vidds™</h1></Link>
            </div>
            <div className="header-right">
                {
                    signedIn ?
                        <LoggedIn />
                        :
                        <LoggedOut />
                }
            </div>
        </header>
    )
}
export default Header