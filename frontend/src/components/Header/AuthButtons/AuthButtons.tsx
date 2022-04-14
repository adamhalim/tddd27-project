import { useAuth0 } from "@auth0/auth0-react"
import LoginButton from "./LoginButton";
import LogoutButton from "./LogoutButton";

const AuthButtons = () => {
    const { isAuthenticated } = useAuth0();

    return isAuthenticated ? <LogoutButton /> : <LoginButton />
}

export default AuthButtons