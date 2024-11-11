// src/components/Navbar.jsx
import { Link } from 'react-router-dom';
import logo from '../assets/images/logo.png';
import './Navbar.css';

function Navbar({ onLogout }) {
  return (
    <nav className="navbar">
      <div className="navbar-logo">
        <Link to="/home"> {/* Puedes reemplazar este Link con una etiqueta de imagen si tienes un logo gráfico */}
          <img src={logo}  alt="Logo" className="logo-image" /> {/* Cambia el src al path de tu logo */}
        </Link>
      </div>
      <div className="navbar-links">
        <Link to="/home">Home</Link>
        <Link to="/cv">Crear Hoja de Vida</Link>
        <Link to="/view-cv">Mis Hojas de Vida</Link>
      </div>
    </nav>
  );
}

export default Navbar;
