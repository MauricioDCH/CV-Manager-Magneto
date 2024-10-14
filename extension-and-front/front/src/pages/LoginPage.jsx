import '../App.css';
import LoginForm from '../components/LoginForm';
import RegisterForm from '../components/RegisterForm';
import HomePage from './HomePage';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

export default function LoginPage({ user, setUser }) {
  const [isRegistering, setIsRegistering] = useState(false);
  const navigate = useNavigate();

  const handleToggleRegister = () => {
    setIsRegistering(!isRegistering);
    navigate(isRegistering ? '/login' : '/register');
  };

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
              ? <p>¿No tienes una cuenta? <button onClick={handleToggleRegister}>Regístrate aquí</button></p>
              : <p>¿Ya tienes una cuenta? <button onClick={handleToggleRegister}>Inicia sesión aquí</button></p>
          )
        }
      </div>
    </div>
  );
}
