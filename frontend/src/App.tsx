import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Header from './components/Header';
import Home from './pages/Home';
import NotFound from './pages/NotFound';
function App() {
  return (
    <Router>
      <Header />
      <Routes>
        <Route path='/' element={<Home />}/>

        <Route path='*' element={<NotFound/>}/>
      </Routes>
    </Router>
  );
}

export default App;
