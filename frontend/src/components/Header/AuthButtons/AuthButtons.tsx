import { useAuth0 } from "@auth0/auth0-react"
import LoginButton from "./LoginButton";
import LogoutButton from "./LogoutButton";

type AuthButtonsType = {
    isAuthenticated: boolean,
}

const AuthButtons = ({isAuthenticated}: AuthButtonsType) => {

    return isAuthenticated ? <LogoutButton /> : <LoginButton />
}

export default AuthButtons