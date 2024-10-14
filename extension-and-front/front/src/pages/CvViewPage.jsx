import React, { useEffect, useState } from 'react';
import { Navigate } from 'react-router-dom';
import { jwtDecode } from 'jwt-decode';

const CvViewPage = ({ user }) => {
    const [cvList, setCvList] = useState([]);
    const [selectedCv, setSelectedCv] = useState(null);
    const [error, setError] = useState('');
    const [loading, setLoading] = useState(true);

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
                const response = await fetch(`http://localhost:8080/cv/user/${userId}`);
                if (response.ok) {
                    const data = await response.json();
                    console.log("response: ", data);
                    setCvList(data);
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

    if (loading) {
        return <p>Cargando...</p>;
    }

    if (error) {
        return <p>{error}</p>;
    }

    const handleCvSelect = (cv) => {
        setSelectedCv(cv);
    };

    return (
        <div>
            <h2>Hojas de Vida</h2>
            {cvList.length > 0 ? (
                <div>
                    <ul>
                        {cvList.map((cv) => (
                            <li key={cv.id}>
                                <button onClick={() => handleCvSelect(cv)} style={{
                                    backgroundColor: '#007bff',  // Color de fondo azul
                                    color: 'white',               // Color del texto blanco
                                    padding: '10px 20px',         // Espaciado interno
                                    margin: '5px 0',              // Espaciado entre botones
                                    border: 'none',               // Sin borde
                                    borderRadius: '5px',          // Bordes redondeados
                                    cursor: 'pointer',            // Cambiar cursor al pasar el mouse
                                    fontSize: '16px',             // Tamaño de fuente
                                    boxShadow: '0px 4px 6px rgba(0, 0, 0, 0.1)', // Sombra
                                }}
                                    onMouseOver={(e) => e.currentTarget.style.backgroundColor = '#0056b3'} // Cambia de color al pasar el mouse
                                    onMouseOut={(e) => e.currentTarget.style.backgroundColor = '#007bff'} // Vuelve al color original al quitar el mouse
                                >
                                    {cv.name} {cv.last_name}
                                </button>
                            </li>
                        ))}
                    </ul>
                    {selectedCv && (
                        <div>
                            <h3>Hoja de Vida de {selectedCv.name} {selectedCv.last_name}</h3>
                            <p><strong>Correo:</strong> {selectedCv.email}</p>
                            <p><strong>Teléfono:</strong> {selectedCv.phone}</p>
                            <p><strong>Experiencia:</strong> {selectedCv.experience}</p>
                            <p><strong>Habilidades:</strong> {selectedCv.skills}</p>
                            <p><strong>Idiomas:</strong> {selectedCv.languages}</p>
                            <p><strong>Educación:</strong> {selectedCv.education}</p>
                        </div>
                    )}
                </div>
            ) : (
                <p>No hay hojas de vida disponibles.</p>
            )}
        </div>
    );
};

export default CvViewPage;
