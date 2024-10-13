// src/pages/CvFormPage.jsx
import React from 'react';
import CvForm from '../components/CvForm';
import { Link, Navigate } from 'react-router-dom';

const CvFormPage = ({ user, setUser }) => {
  const handleLogout = () => {
    setUser(null);
    localStorage.removeItem('userInfo'); // Limpia el usuario del localStorage
  };
  

  // Si el usuario no está logueado, redirige a la página de inicio de sesión
  if (!user) {
    return <Navigate to="/login" replace />;
  }

  return (
    <div>
      <h1>Bienvenido, {user.name}</h1>
      <CvForm />
      <button onClick={handleLogout}>Cerrar Sesión</button>
    </div>
  );
};

export default CvFormPage;
