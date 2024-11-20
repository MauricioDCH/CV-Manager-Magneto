// src/pages/HomePage.jsx
import React, { useEffect, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { jwtDecode } from 'jwt-decode';

const HomePage = ({ user, setUser }) => {
  const [cvList, setCvList] = useState([]);
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();
  
  const getUserIdFromToken = () => {
    const token = localStorage.getItem('token');
    if (!token) {
      return null;
    }
    try {
      const decodedToken = jwtDecode(token);
      return decodedToken.sub; // Obtener el userId del campo 'sub' del token
    } catch (error) {
      console.error('Error al decodificar el token:', error);
      return null;
    }
  };

  useEffect(() => {
    const fetchCvData = async () => {
      const userId = getUserIdFromToken();
      if (!userId) {
        setError('Usuario no autenticado');
        setLoading(false);
        return;
      }

      try {
        //const response = await fetch(`http://cv:8008/cv/user/${userId}`);
        const response = await fetch(`http://localhost:8008/cv/user/${userId}`);
        if (response.ok) {
          const data = await response.json();
          setCvList(data);
          console.log(data)
        } else {
          setError('No se pudo obtener la hoja de vida.');
        }
      } catch (err) {
        console.error('Error al obtener la hoja de vida:', err);
        setError('Error al conectar con el servidor. Inténtalo más tarde.');
      } finally {
        setLoading(false);
      }
    };

    fetchCvData();
  }, []);

  const handleLogout = () => {
    setUser(null);
    localStorage.removeItem('userInfo'); // Limpia el usuario del localStorage
  };

  const handleViewCv = (cvId) => {
    navigate(`/view-cv`);
  };

  const handleEditCv = (cvId) => {
    navigate(`/edit-cv/${cvId}`);
  };

  return (
   <div>
      {user ? (
        <>
          <h1>Bienvenido, {user.name}</h1>

          {loading ? (
            <p>Cargando tus hojas de vida...</p> // Mensaje de carga
          ) : (
            <>
              {cvList.length > 0 ? (
                <div className="cv-cards">
                  {cvList.map((cv) => (
                    <div key={cv.id} className="cv-card">
                      <h3>{cv.title}</h3>
                      <p><strong>Nombre:</strong> {cv.name} {cv.last_name}</p>
                      <p><strong>Experiencia:</strong> {cv.experience}</p>
                      <p><strong>Habilidades:</strong> {cv.skills}</p>

                      <div className="cv-card-buttons">
                      </div>
                    </div>
                  ))}
                </div>
              ) : (
                <p>No tienes hojas de vida disponibles.</p>
              )}
            </>
          )}

          <button className="logout-button" onClick={handleLogout} style={{ marginTop: '20px' }}>Cerrar Sesión</button>
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
