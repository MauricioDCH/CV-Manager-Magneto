import './App.css';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import HomePage from './pages/HomePage';
import LoginPage from './pages/LoginPage';
import RegisterPage from './pages/RegisterPage';
import CvFormPage from './pages/CvFormPage';
import CvViewPage from './pages/CvViewPage';
import { useState, useEffect } from 'react';

function App() {
  const [user, setUser] = useState(() => {
    const userInfo = localStorage.getItem('userInfo');
    return userInfo ? JSON.parse(userInfo) : null; // Cargar usuario desde localStorage
  });

  const [isRegistering, setIsRegistering] = useState(false);

  return (
    <div>
      <BrowserRouter>
        <Routes>
          <Route index element={<HomePage user={user} setUser={setUser} />} />
          <Route path='/home' element={<HomePage user={user} setUser={setUser} />} />
          <Route path='/register' element={<RegisterPage user={user} setUser={setUser} />} />
          <Route path='/login' element={<LoginPage user={user} setUser={setUser} />} />
          <Route path='/cv' element={<CvFormPage user={user} setUser={setUser} />} />
          <Route path="/view-cv" element={<CvViewPage user={user} />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
