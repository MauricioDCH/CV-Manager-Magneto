document.addEventListener('DOMContentLoaded', () => {

  const redirectButton = document.getElementById('redirectButton');
  const fillFormButton = document.getElementById('sendFieldsButton');
  const cvSelect = document.getElementById('cvSelect');

  // Obtener el token desde chrome.storage.local
  const getToken = () => {
    return new Promise((resolve) => {
      chrome.storage.local.get('token', (result) => {
        if (result.token) {
          console.log("Token JWT encontrado en chrome.storage.local:", result.token);
          resolve(result.token);
        } else {
          console.error('No se pudo obtener el token');
          resolve(null);
        }
      });
    });
  };

  function decodeJWT(token) {
    const base64Url = token.split('.')[1];
    const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    const jsonPayload = decodeURIComponent(atob(base64).split('').map(function (c) {
      return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
  }

  // Función para obtener el userId desde el token
  const getUserIdFromToken = (token) => {
    if (!token) return null;
    try {
      const decodedToken = decodeJWT(token); // Asegúrate de tener jwt-decode como dependencia
      console.log(decodedToken);
      console.log(decodedToken.sub)
      return decodedToken.sub; // Obtener el userId del campo 'sub' del token
    } catch (error) {
      console.error('Error al decodificar el token:', error);
      return null;
    }
  };

  // Cargar las opciones de hojas de vida en el menú desplegable
  const loadCvOptions = async () => {
    const token = await getToken();
    if (!token) return;

    const userId = getUserIdFromToken(token);
    if (!userId) return;

    try {
      const response = await fetch(`http://34.27.58.251:8008/cv/user/${userId}`);
      const cvs = await response.json();

      cvs.forEach(cv => {
        const option = document.createElement('option');
        option.value = cv.id;
        option.textContent = cv.title;
        cvSelect.appendChild(option);
      });
    } catch (error) {
      console.error('Error al cargar las hojas de vida:', error);
    }
    // Almacenar el ID de la hoja de vida seleccionada
    cvSelect.addEventListener('change', (event) => {
      const selectedCvId = event.target.value;
      chrome.storage.local.set({ selectedCvId: selectedCvId }, () => {
        console.log('ID de la hoja de vida seleccionada guardado:', selectedCvId);
      });
    });
  };

  loadCvOptions();

  // Manejar clic en el botón de redirección
  if (redirectButton) {
    redirectButton.addEventListener('click', () => {
      const cvManagerUrl  = 'http://34.27.58.251:80'; // URL de tu página de login
      chrome.tabs.create({ url: cvManagerUrl });
    });
  } else {
    console.error('El botón de redirección no se encontró en el DOM.');
  }
  // Manejar clic en el botón de rellenar formulario
  if (fillFormButton) {
    fillFormButton.addEventListener('click', () => {
      chrome.tabs.query({ active: true, currentWindow: true }, (tabs) => {
        chrome.scripting.executeScript({
          target: { tabId: tabs[0].id },
          function: () => {
            window.postMessage({ type: 'FILL_FORM' }, '*');
          }
        });
      });
    });
  } else {
    console.error('El botón de rellenar formulario no se encontró en el DOM.');
  }
});