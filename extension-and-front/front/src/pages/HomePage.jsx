import React from 'react'

const HomePage = ({ user, setUser }) => {

  const handleLogout = () => {
    setUser(null) // Cambiado a null para representar que no hay usuario logueado
  }

  return (
    <div>
      <h1>Bienvenido</h1>
      <h2>{user.name}</h2> {/* Muestra el nombre del usuario */}
      <button onClick={handleLogout}>Cerrar Sesión</button>
    </div>
  )
}

export default HomePage
