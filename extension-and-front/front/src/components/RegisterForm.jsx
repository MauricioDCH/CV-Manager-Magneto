import React, { useState } from 'react';
import './LoginForm.css';
import { validatePassword, validateEmail } from '../utils/validation'; // Importa las funciones de validación

const RegisterForm = ({ setUser }) => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [passwordError, setPasswordError] = useState('');
  const [emailError, setEmailError] = useState(''); // Nuevo estado para errores de email
  const [error, setError] = useState('');

   // Manejar el cambio de la contraseña para validación en tiempo real
  const handlePasswordChange = (e) => {
    const newPassword = e.target.value;
    setPassword(newPassword);

    const errorMessage = validatePassword(newPassword);
    setPasswordError(errorMessage); // Muestra el error en tiempo real
  };

  // Manejar el cambio del correo para validación en tiempo real
  const handleEmailChange = (e) => {
    const newEmail = e.target.value;
    setEmail(newEmail);

    const errorMessage = validateEmail(newEmail);
    setEmailError(errorMessage); // Muestra el error en tiempo real
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (name === '' || email === '' || password === '') {
      setError('Todos los campos son obligatorios');
      return;
    }

    if (passwordError || emailError) {
      setError(passwordError || emailError); // Mostrar error si existe alguna validación pendiente
      return;
    }

    setError('');

    try {
      const newUser = { name, email, password };
      const response = await fetch('http://localhost:8080/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(newUser),
      });

      if (response.ok) {
        setUser(newUser); // Aquí estableces el usuario registrado
      } else {
        setError('Error al registrar el usuario');
      }
    } catch (error) {
      console.error('Error al registrar el usuario:', error);
      setError('Ocurrió un error al intentar registrar. Por favor, intenta de nuevo');
    }
  };

  return (
    <section>
      <h1>Registro</h1>
      <form className="form" onSubmit={handleSubmit}>
        <input
          type="text"
          value={name}
          placeholder="Nombre"
          onChange={(e) => setName(e.target.value)}
        />
        <input
          type="email"
          value={email}
          placeholder="Correo"
          onChange={handleEmailChange} // Llama a la función de validación en tiempo real
        />
        {emailError && <p style={{ color: 'red' }}>{emailError}</p>} {/* Muestra el error de email en tiempo real */}
        <input
          type="password"
          value={password}
          placeholder="Contraseña"
          onChange={handlePasswordChange}
        />
        {passwordError && <p style={{ color: 'red' }}>{passwordError}</p>} {/* Muestra el error en tiempo real */}
        <button type="submit">Registrarse</button>
      </form>
      {error && <p>{error}</p>}
    </section>
  );
};

export default RegisterForm;
