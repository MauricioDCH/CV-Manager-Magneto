import '@fortawesome/fontawesome-free/css/all.min.css';
import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { jwtDecode } from 'jwt-decode';
import './CvViewPage.css'

const CvViewPage = () => {
    const [cvList, setCvList] = useState([]);
    const [expandedCvId, setExpandedCvId] = useState(null);
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
            return decodedToken.sub;
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
                //const response = await fetch(`http://localhost:8008/cv/user/${userId}`);
                const response = await fetch(`http://34.45.83.31:8008/cv/user/${userId}`);
                if (response.ok) {
                    const data = await response.json();
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

    const toggleExpandCv = (cvId) => {
        setExpandedCvId(expandedCvId === cvId ? null : cvId);
    };

    const handleEditCv = (cvId) => {
        navigate(`/edit-cv/${cvId}`);
    };

    const handleDeleteCv = async (cvId) => {
        const confirmDelete = window.confirm("¿Estás seguro de que deseas eliminar esta hoja de vida?");
        if (confirmDelete) {
            try {
                //const response = await fetch(`http://cv:8008/cv/${cvId}`, { method: 'DELETE' });
                //const response = await fetch(`http://localhost:8008/cv/${cvId}`, { method: 'DELETE' });
                const response = await fetch(`http://34.45.83.31:8008/cv/${cvId}`, { method: 'DELETE' });
                if (response.ok) {
                    setCvList(cvList.filter(cv => cv.id !== cvId));
                    if (expandedCvId === cvId) setExpandedCvId(null);
                } else {
                    setError('No se pudo eliminar la hoja de vida.');
                }
            } catch (err) {
                console.error('Error al eliminar la hoja de vida:', err);
                setError('Error al conectar con el servidor. Inténtalo más tarde.');
            }
        }
    };

    if (loading) {
        return <p>Cargando...</p>;
    }

    if (error) {
        return <p>{error}</p>;
    }

    return (
        <div>
            <h2>Hojas de Vida</h2>
            {cvList.length > 0 ? (
                <ul style={{ listStyle: 'none', padding: 0 }}>
                    {cvList.map((cv) => (
                        <li key={cv.id} style={{ marginBottom: '20px' }}>
                            <div
                                onClick={() => toggleExpandCv(cv.id)}
                                onMouseEnter={(e) => e.currentTarget.style.backgroundColor = '#00a571'}
                                onMouseLeave={(e) => e.currentTarget.style.backgroundColor = '#00D591'}
                                style={{
                                    backgroundColor: '#00D591', color: 'white', padding: '10px 20px',
                                    cursor: 'pointer', borderRadius: '5px',
                                    fontWeight: expandedCvId === cv.id ? 'bold' : 'normal'
                                }}
                            >
                                {cv.title} - {cv.name} {cv.last_name}
                            </div>

                            {expandedCvId === cv.id && (
                                <div className="cv-details">
                                    <p><strong>Nombre:</strong> {cv.name}</p>
                                    <p><strong>Apellidos:</strong> {cv.last_name}</p>
                                    <p><strong>Correo:</strong> {cv.email}</p>
                                    <p><strong>Teléfono:</strong> {cv.phone}</p>
                                    <p><strong>Experiencia:</strong> {cv.experience}</p>
                                    <p><strong>Habilidades:</strong> {cv.skills}</p>
                                    <p><strong>Idiomas:</strong> {cv.languages}</p>
                                    <p><strong>Educación:</strong> {cv.education}</p>
                                    <div className="cv-buttons">
                                        <button className="cv-button edit" onClick={() => handleEditCv(cv.id)}>
                                            <i className="fas fa-edit"></i> {/* Ícono de editar */}
                                        </button>
                                        <button className="cv-button delete" onClick={() => handleDeleteCv(cv.id)}>
                                            <i className="fas fa-trash"></i> {/* Ícono de eliminar */}
                                        </button>
                                    </div>
                                </div>
                            )}
                        </li>
                    ))}
                </ul>
            ) : (
                <p>No hay hojas de vida disponibles.</p>
            )}
            <button className="back-button" onClick={() => navigate(-1)} style={{ marginTop: '20px', padding: '10px 20px', borderRadius: '5px' }}>
                Atrás
            </button>
        </div>
    );
};

export default CvViewPage;
