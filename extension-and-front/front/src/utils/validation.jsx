export const validatePassword = (password) => {
    const minLength = /.{8,}/;
    const hasUppercase = /[A-Z]/;
    const hasNumber = /[0-9]/;
    const hasSpecialChar = /[!@#$%^&*(),.?":{}|<>]/;
  
    if (!minLength.test(password)) {
      return 'La contraseña debe tener al menos 8 caracteres';
    }
    if (!hasUppercase.test(password)) {
      return 'La contraseña debe tener al menos una letra mayúscula';
    }
    if (!hasNumber.test(password)) {
      return 'La contraseña debe tener al menos un número';
    }
    if (!hasSpecialChar.test(password)) {
      return 'La contraseña debe tener al menos un carácter especial';
    }
    return null; // No hay errores
  };
  
  export const validateEmail = (email) => {
    const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailPattern.test(email) ? null : 'Por favor, introduce un correo electrónico válido';
  };
  