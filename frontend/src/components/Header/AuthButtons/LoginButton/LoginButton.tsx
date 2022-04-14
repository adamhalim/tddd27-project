import { useAuth0 } from "@auth0/auth0-react";
import '../style.css'

const LoginButton = () => {
    const { loginWithRedirect } = useAuth0();

    return (
        <div className="btn-header" onClick={() => loginWithRedirect()} >
            <button>Log In</button>
        </div>
    )
};

export default LoginButton;