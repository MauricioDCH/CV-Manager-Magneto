// src/pages/CvFormPage.jsx
import React from 'react';
import CvForm from '../components/CvForm';
import {Navigate, useNavigate } from 'react-router-dom';
import './CvFormPage.css';

const CvFormPage = ({ user, setUser }) => {
  const navigate = useNavigate();
  
  // Si el usuario no está logueado, redirige a la página de inicio de sesión
  if (!user) {
    return <Navigate to="/login" replace />;
  }

  return (
    <div>
      <CvForm />
      <button className="back-button" onClick={() => navigate(-1)}>Atrás</button>
    </div>
  );
};

export default CvFormPage;
