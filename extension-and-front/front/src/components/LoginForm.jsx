import React, { useState } from 'react'
import "./LoginForm.css"

const LoginForm = ({ setUser }) => {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [error, setError] = useState("")

  const handleSubmit = async (e) => {
    e.preventDefault()

    if (email === "" || password === "") {
      setError("Todos los campos son obligatorios")
      return
    }

    setError("")

      // Make the POST request with the email and password in the body
      try {
        const newUser = { email, password }
        const response = await fetch('http://localhost:8000/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(newUser)
        })
        const data = await response.json()

        if (Object.keys(data).length > 0) {
          // Guarda el nombre y el email en el estado
          setUser({ name: data.name, email: data.email })
  
          localStorage.setItem('userInfo', JSON.stringify({ name: data.name, email: data.email }));
          
          const sendMessageToContentScript = (userInfo) => {
            window.postMessage({ type: 'FROM_PAGE', userInfo }, '*');
        };
        
        if (window.chrome) {
            sendMessageToContentScript({ name: data.name, email: data.email });
        }
          
        } else {
          setError("El correo o la contraseña ingresados no se encuentran registrados")
        }
      } catch (error) {
        console.error("Error al iniciar sesión:", error)
        setError("El correo o la contraseña ingresados no se encuentran registrados")
      }
    }

  return (
    <section>
      <h1>Login</h1>
      <form className='form' onSubmit={handleSubmit}>
        <input
          type="email"
          value={email}
          placeholder='Correo'
          onChange={e => setEmail(e.target.value)}
        />
        <input
          type="password"
          value={password}
          placeholder='Constraseña'
          onChange={e => setPassword(e.target.value)}
        />
        <button>Iniciar Sesión</button>
      </form>
      {error && <p>{error}</p>}
    </section>
  )
}

export default LoginForm