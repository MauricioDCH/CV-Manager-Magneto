import './App.css'
import LoginForm from './components/LoginForm'
import RegisterForm from './components/RegisterForm'
import HomePage from './pages/HomePage'
import { useState } from 'react'

function App() {
  const [user, setUser] = useState(null)
  const [isRegistering, setIsRegistering] = useState(false)

  return (
    <div>
      {
        !user
          ? (
            isRegistering
              ? <RegisterForm setUser={setUser} />
              : <LoginForm setUser={setUser} />
          )
          : <HomePage user={user} setUser={setUser} />
      }
      <div>
        {
          !user && (
            !isRegistering
              ? <p>¿No tienes una cuenta? <button onClick={() => setIsRegistering(true)}>Regístrate aquí</button></p>
              : <p>¿Ya tienes una cuenta? <button onClick={() => setIsRegistering(false)}>Inicia sesión aquí</button></p>
          )
        }
      </div>
    </div>
  )
}

export default App
