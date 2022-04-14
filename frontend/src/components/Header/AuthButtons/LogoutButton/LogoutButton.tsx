import { useAuth0 } from "@auth0/auth0-react";

const LogoutButton = () => {
    const { logout } = useAuth0();

    return (
        <div className="btn-header" onClick={() => logout({ returnTo: window.location.origin })}>
            <button >
                Log Out
            </button>
        </div>
    );
};

export default LogoutButton;