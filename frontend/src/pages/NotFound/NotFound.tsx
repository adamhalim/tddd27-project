import './style.css'

const NotFound = () => {
    return (
        <div className='container'>
            <div className='err-container'>
                <div className='err-wrapper'>
                    <h1>404</h1>
                    <p>Error</p>
                </div>
            </div>
            <h2>Oopsie, the page you were looking could not be found :(</h2>
            <p></p>
        </div>
    )
}

export default NotFound