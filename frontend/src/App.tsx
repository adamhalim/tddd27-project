import { useState } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Header from './components/Header';
import Home from './pages/Home';
import NotFound from './pages/NotFound';

function App() {

  const [signedIn, setSignedIn] = useState(false);

  const signIn = (): void => {
    setSignedIn(true)
  }

  return (
    <Router>
      <Header signedIn={signedIn} signIn={signIn}/>
      <Routes>
        <Route path='/' element={<Home />}/>

        <Route path='*' element={<NotFound/>}/>
      </Routes>
    </Router>
  );
}

export default App;
