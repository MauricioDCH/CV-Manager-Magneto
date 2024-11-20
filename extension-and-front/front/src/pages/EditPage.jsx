// src/pages/CvEditPage.jsx
import React, { useEffect, useState } from 'react';
import { useParams, Navigate, useNavigate } from 'react-router-dom';
import { jwtDecode } from 'jwt-decode';

const CvEditPage = ({ user }) => {
  const { cvId } = useParams(); // Obtener el ID de la hoja de vida desde la URL
  const navigate = useNavigate();
  const [cvData, setCvData] = useState({
    title: '',
    name: '',
    last_name: '',
    email: '',
    phone: '',
    experience: '',
    skills: '',
    languages: '',
    education: '',
  });
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(true);
  const [successMessage, setSuccessMessage] = useState(''); // Para mostrar mensaje de éxito

  // Si el usuario no está logueado, redirige a la página de inicio de sesión
  if (!user) {
    return <Navigate to="/login" replace />;
  }

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
      try {
        const userId = getUserIdFromToken();
        // Obtén todas las hojas de vida del usuario
        //const response = await fetch(`http://cv:8008/cv/user/${userId}`);
        //const response = await fetch(`http://localhost:8008/cv/user/${userId}`);
        const response = await fetch(`http://34.45.83.31:8008/cv/user/${userId}`);

        if (response.ok) {
          const data = await response.json();
          // Filtrar la hoja de vida específica
          const selectedCv = data.find(cv => cv.id === parseInt(cvId));
          if (selectedCv) {
            setCvData(selectedCv); // Asigna la hoja de vida específica al estado
          } else {
            setError('Hoja de vida no encontrada.');
          }
        } else {
          setError('No se pudo obtener las hojas de vida.');
        }
      } catch (err) {
        console.error('Error al obtener las hojas de vida:', err);
        setError('Error al conectar con el servidor. Inténtalo más tarde.');
      } finally {
        setLoading(false);
      }
    };

    fetchCvData();
  }, [cvId]);

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Realizar la petición PUT para actualizar la hoja de vida
    try {
      //const response = await fetch(`http://cv:8008/cv/${cvId}`, {
      //const response = await fetch(`http://localhost:8008/cv/${cvId}`, {
      const response = await fetch(`http://34.45.83.31:8008/cv/${cvId}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(cvData),
      });

      if (response.ok) {
        navigate('/view-cv'); // Redirigir después de actualizar
      } else {
        setError('No se pudo actualizar la hoja de vida.');
      }
    } catch (err) {
      console.error('Error al actualizar la hoja de vida:', err);
      setError('Error al conectar con el servidor. Inténtalo más tarde.');
    }
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setCvData({ ...cvData, [name]: value });
  };

  if (loading) {
    return <p>Cargando...</p>;
  }

  if (error) {
    return <p>{error}</p>;
  }

  return (
    <div>      
        <h2>Edita tu hoja de vida</h2>
        <div className="cv-edit-form">
        {successMessage && <p style={{ color: 'green' }}>{successMessage}</p>} {/* Mensaje de éxito */}
        <form onSubmit={handleSubmit}>
          <div>
            <label>Título:</label>
            <input
              type="text"
              name="title"
              value={cvData.title}
              onChange={handleChange}
              required
            />
          </div>
          <div>
            <label>Nombre:</label>
            <input
              type="text"
              name="name"
              value={cvData.name}
              onChange={handleChange}
              required
            />
          </div>
          <div>
            <label>Apellido:</label>
            <input
              type="text"
              name="last_name"
              value={cvData.last_name}
              onChange={handleChange}
              required
            />
          </div>
          <div>
            <label>Correo:</label>
            <input
              type="email"
              name="email"
              value={cvData.email}
              onChange={handleChange}
              required
            />
          </div>
          <div>
            <label>Teléfono:</label>
            <input
              type="tel"
              name="phone"
              value={cvData.phone}
              onChange={handleChange}
              required
            />
          </div>
          <div>
            <label>Experiencia:</label>
            <textarea
              name="experience"
              value={cvData.experience}
              onChange={handleChange}
              required
            />
          </div>
          <div>
            <label>Habilidades:</label>
            <textarea
              name="skills"
              value={cvData.skills}
              onChange={handleChange}
              required
            />
          </div>
          <div>
            <label>Idiomas:</label>
            <textarea
              name="languages"
              value={cvData.languages}
              onChange={handleChange}
              required
            />
          </div>
          <div>
            <label>Educación:</label>
            <textarea
              name="education"
              value={cvData.education}
              onChange={handleChange}
              required
            />
          </div>
          <button type="submit">Actualizar Hoja de Vida</button>
        </form>
        </div>
        <button className="back-button" onClick={() => navigate(-1)} style={{ marginTop: '10px' }}>Atrás</button>
      </div>
  );
};

export default CvEditPage;
