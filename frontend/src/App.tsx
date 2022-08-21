import { Auth0Provider } from '@auth0/auth0-react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Header from './components/Header';
import Home from './pages/Home';
import NotFound from './pages/NotFound';
import Profile from './pages/Profile';
import VideoPage from './pages/VideoPage';

function App() {

  return (
    <Router>
      <Auth0Provider
        domain={`${process.env.REACT_APP_DOMAIN}`}
        clientId={`${process.env.REACT_APP_CLIENT_ID}`}
        redirectUri={window.location.origin}
      >
        <Header />
        <Routes>
          <Route path='/' element={<Home />} />
          <Route path='/profile' element={<Profile />} />
          <Route path='/video/:id' element={<VideoPage />}/>
          <Route path='*' element={<NotFound />} />
        </Routes>
      </Auth0Provider>
    </Router>
  );
}

export default App;
