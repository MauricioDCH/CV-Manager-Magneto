import '../App.css';
import LoginForm from '../components/LoginForm';
import RegisterForm from '../components/RegisterForm';
import HomePage from './HomePage';
import { useState } from 'react';

export default function LoginPage({ user, setUser }) { // Ahora recibe user y setUser
  const [isRegistering, setIsRegistering] = useState(false);

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
  );
}
