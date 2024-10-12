// src/pages/HomePage.jsx
import React from 'react';
import CvForm from '../components/CvForm';

const HomePage = ({ user, setUser }) => {
  const handleLogout = () => {
    setUser(null); // Cambiado a null para representar que no hay usuario logueado
  };

  return (
    <div>
      <h1>Bienvenido, {user.name}</h1>
      <CvForm /> {/* Aquí agregamos el formulario */}
      <button onClick={handleLogout}>Cerrar Sesión</button>
    </div>
  );
};

export default HomePage;
