import React, { useEffect, useState } from 'react';
import { Navigate, useNavigate } from 'react-router-dom';
import { jwtDecode } from 'jwt-decode';

const CvViewPage = ({ user }) => {
    const [cvList, setCvList] = useState([]);
    const [selectedCv, setSelectedCv] = useState(null);
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

    if (loading) {
        return <p>Cargando...</p>;
    }

    if (error) {
        return <p>{error}</p>;
    }

    const handleCvSelect = (cv) => {
        setSelectedCv(cv);
    };

    const handleEditCv = () => {
        if (selectedCv) {
            navigate(`/edit-cv/${selectedCv.id}`); // Navegar a la página de edición
        }
    };

    const handleDeleteCv = async () => {
        if (selectedCv) {
            const confirmDelete = window.confirm("¿Estás seguro de que deseas eliminar esta hoja de vida?");
            if (confirmDelete) {
                try {
                    const response = await fetch(`http://localhost:8008/cv/${selectedCv.id}`, {
                        method: 'DELETE',
                    });
                    if (response.ok) {
                        // Actualizar la lista de hojas de vida después de eliminar exitosamente
                        setCvList(cvList.filter(cv => cv.id !== selectedCv.id));
                        setSelectedCv(null);
                    } else {
                        setError('No se pudo eliminar la hoja de vida.');
                    }
                } catch (err) {
                    console.error('Error al eliminar la hoja de vida:', err);
                    setError('Error al conectar con el servidor. Inténtalo más tarde.');
                }
            }
        }
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
                                    backgroundColor: '#007bff',
                                    color: 'white',
                                    padding: '10px 20px',
                                    margin: '5px 0',
                                    border: 'none',
                                    borderRadius: '5px',
                                    cursor: 'pointer',
                                    fontSize: '16px',
                                    boxShadow: '0px 4px 6px rgba(0, 0, 0, 0.1)',
                                }}
                                    onMouseOver={(e) => e.currentTarget.style.backgroundColor = '#0056b3'}
                                    onMouseOut={(e) => e.currentTarget.style.backgroundColor = '#007bff'}
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
                            <button onClick={handleEditCv} style={{
                                backgroundColor: '#28a745', // Color de fondo verde
                                color: 'white', // Color del texto blanco
                                padding: '10px 20px', // Espaciado interno
                                margin: '10px 0', // Espaciado entre botones
                                border: 'none', // Sin borde
                                borderRadius: '5px', // Bordes redondeados
                                cursor: 'pointer', // Cambiar cursor al pasar el mouse
                                fontSize: '16px', // Tamaño de fuente
                                boxShadow: '0px 4px 6px rgba(0, 0, 0, 0.1)', // Sombra
                            }}>
                                Editar Hoja de Vida
                            </button>

                            <button onClick={handleDeleteCv} style={{
                                backgroundColor: '#dc3545', // Color de fondo rojo
                                color: 'white', // Color del texto blanco
                                padding: '10px 20px', // Espaciado interno
                                margin: '10px', // Espaciado entre botones
                                border: 'none', // Sin borde
                                borderRadius: '5px', // Bordes redondeados
                                cursor: 'pointer', // Cambiar cursor al pasar el mouse
                                fontSize: '16px', // Tamaño de fuente
                                boxShadow: '0px 4px 6px rgba(0, 0, 0, 0.1)', // Sombra
                            }}>
                                Eliminar Hoja de Vida
                            </button>
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
