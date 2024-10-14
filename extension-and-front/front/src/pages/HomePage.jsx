// src/pages/HomePage.jsx
import React from 'react';
import CvForm from '../components/CvForm';
import { Link } from 'react-router-dom';

const HomePage = ({ user, setUser }) => {
  const handleLogout = () => {
    setUser(null);
    localStorage.removeItem('userInfo'); // Limpia el usuario del localStorage
  };
   

  return (
    <div>
      {user ? (
        <>
          <h1>Bienvenido, {user.name}</h1>
          <Link to="/cv">Crear Hoja de vida</Link>
          <Link to="/view-cv">Ver Hoja de Vida</Link>
          <button onClick={handleLogout}>Cerrar Sesión</button>
        </>     
      ) : (
        <>
          <h1>Bienvenido a CV Manager</h1>
          <p>Por favor, regístrate o inicia sesión para continuar.</p>
          <Link to="/login">Iniciar Sesión</Link> | <Link to="/register">Registrarse</Link>
        </>
      )}
    </div>
  );
};

export default HomePage;
