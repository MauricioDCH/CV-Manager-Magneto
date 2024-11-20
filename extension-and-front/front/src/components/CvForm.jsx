import React, { useState } from 'react';
import { jwtDecode } from 'jwt-decode';
import './CvForm.css';

const CvForm = () => {
    const [title, setTitle] = useState('');
    const [name, setName] = useState('');
    const [lastName, setLastName] = useState('');
    const [email, setEmail] = useState('');
    const [phone, setPhone] = useState('');
    const [experience, setExperience] = useState('');
    const [skills, setSkills] = useState('');
    const [languages, setLanguages] = useState('');
    const [education, setEducation] = useState('');
    const [error, setError] = useState('');
    const [success, setSuccess] = useState('');

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

    const handleSubmit = async (e) => {
        e.preventDefault();

        // Validar campos requeridos
        if (!title || !name || !lastName || !email || !phone || !experience || !skills || !languages || !education) {
            setError('Todos los campos son obligatorios');
            setSuccess('');
            return;
        }

        setError('');

        // user_id recuperado del token JWT
        const userId = getUserIdFromToken();
        const intUserId = parseInt(userId, 10);
        console.log("user id desde crear cv:", intUserId)
        const resumeData = { title, name, last_name: lastName, email, phone, experience, skills, languages, education, user_id: intUserId };

        try {
            // Realizar la petición POST al endpoint
            //const response = await fetch('http://create-cv:8081/create-cv', {
            //const response = await fetch('http://localhost:8081/create-cv', {
            const response = await fetch('http://34.27.58.251:8081/create-cv', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(resumeData),
            });

            if (response.ok) {
                const data = await response.json();
                console.log('Hoja de vida enviada:', data);
                setSuccess('Hoja de vida guardada exitosamente!');

                // Reiniciar el formulario
                setTitle('');
                setName('');
                setLastName('');
                setEmail('');
                setPhone('');
                setExperience('');
                setSkills('');
                setLanguages('');
                setEducation('');
            } else {
                setError('Ocurrió un error al guardar la hoja de vida.');
                setSuccess('');
            }
        } catch (err) {
            console.error('Error al enviar la hoja de vida:', err);
            setError('Error al conectar con el servidor. Inténtalo más tarde.');
            setSuccess('');
        }
    };

    return (
        <section className="cv-section">
            <h2>Crea tu hoja de vida</h2>
            <form className='form' onSubmit={handleSubmit}>
                <input type="text" value={title} placeholder="Título de la hoja de vida" onChange={(e) => setTitle(e.target.value)} />
                <input type="text" value={name} placeholder="Nombre" onChange={(e) => setName(e.target.value)} />
                <input type="text" value={lastName} placeholder="Apellido" onChange={(e) => setLastName(e.target.value)} />
                <input type="email" value={email} placeholder="Correo" onChange={(e) => setEmail(e.target.value)} />
                <input type="tel" value={phone} placeholder="Teléfono" onChange={(e) => setPhone(e.target.value)} />
                <textarea value={experience} placeholder="Experiencia" onChange={(e) => setExperience(e.target.value)} />
                <input type="text" value={skills} placeholder="Habilidades (separadas por comas)" onChange={(e) => setSkills(e.target.value)} />
                <input type="text" value={languages} placeholder="Idiomas (separados por comas)" onChange={(e) => setLanguages(e.target.value)} />
                <input type="text" value={education} placeholder="Educación (separados por comas)" onChange={(e) => setEducation(e.target.value)} />
                <button type="submit">Guardar Hoja de Vida</button>
            </form>
            {error && <p className="error-message">{error}</p>}
            {success && <p className="success-message">{success}</p>}
        </section>
    );
};

export default CvForm;