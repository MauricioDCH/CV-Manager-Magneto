// src/components/Navbar.jsx
import { Link } from 'react-router-dom';
import logo from '../assets/images/logo.png';
import './Navbar.css';

function Navbar({ onLogout }) {
  return (
    <nav className="navbar">
      <div className="navbar-logo">
        <Link to="/home"> 
          <img src={logo}  alt="Logo" className="logo-image" /> 
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
