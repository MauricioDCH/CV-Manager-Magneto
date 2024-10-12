import React, { useState } from 'react';
import './LoginForm.css';

const CvForm = () => {
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

    const handleSubmit = async (e) => {
        e.preventDefault();

        // Validar campos requeridos
        if (!name || !lastName || !email || !phone || !experience || !skills || !languages || !education) {
            setError('Todos los campos son obligatorios');
            setSuccess('');
            return;
        }

        setError('');

        // user_id quemado temporalmente
        const userId = 1;  // Cambia el valor según sea necesario

        const resumeData = { name, lastName, email, phone, experience, skills, languages, education, user_id: userId };

        try {
            // Realizar la petición POST al endpoint
            const response = await fetch('http://localhost:8081/create-cv', {
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
        <section>
            <h2>Crea tu Hoja de Vida</h2>
            <form className='form' onSubmit={handleSubmit}>
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
            {error && <p>{error}</p>}
            {success && <p>{success}</p>}
        </section>
    );
};

export default CvForm;
