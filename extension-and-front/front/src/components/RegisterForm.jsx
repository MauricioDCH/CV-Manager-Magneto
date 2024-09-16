import React, { useState } from 'react'
import "./LoginForm.css"

const RegisterForm = ({ setUser }) => {
  const [name, setName] = useState("")
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [error, setError] = useState("")

  const handleSubmit = async (e) => {
    e.preventDefault()

    if (name === "" || email === "" || password === "") {
      setError("Todos los campos son obligatorios")
      return
    }

    setError("")

    try {
      const newUser = { name, email, password }
      const response = await fetch('http://localhost:8080/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(newUser)
      })

      if (response.ok) {
        setUser(newUser) // Aquí estableces el usuario registrado
      } else {
        setError("Error al registrar el usuario")
      }
    } catch (error) {
      console.error("Error al registrar el usuario:", error)
      setError("Ocurrió un error al intentar registrar. Por favor, intenta de nuevo")
    }
  }

  return (
    <section>
      <h1>Registro</h1>
      <form className='form' onSubmit={handleSubmit}>
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
          onChange={(e) => setEmail(e.target.value)}
        />
        <input
          type="password"
          value={password}
          placeholder="Contraseña"
          onChange={(e) => setPassword(e.target.value)}
        />
        <button type="submit">Registrarse</button>
      </form>
      {error && <p>{error}</p>}
    </section>
  )
}

export default RegisterForm
