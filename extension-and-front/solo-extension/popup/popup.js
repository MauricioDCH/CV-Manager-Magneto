document.addEventListener('DOMContentLoaded', () => {
  const redirectButton = document.getElementById('redirectButton');
  const sendFieldsButton = document.getElementById('sendFieldsButton');

  // Manejar clic en el botón de redirección
  if (redirectButton) {
    redirectButton.addEventListener('click', () => {
      const loginUrl = 'http://localhost:5173'; // URL de tu página de login
      chrome.tabs.query({ active: true, currentWindow: true }, (tabs) => {
        chrome.tabs.update(tabs[0].id, { url: loginUrl });
      });
    });
  } else {
    console.error('El botón de redirección no se encontró en el DOM.');
  }

  // Manejar clic en el botón de enviar campos
  if (sendFieldsButton) {
    console.log("boton si sirve")
    sendFieldsButton.addEventListener('click', () => {
      chrome.tabs.query({ active: true, currentWindow: true }, (tabs) => {
        chrome.scripting.executeScript({
          target: { tabId: tabs[0].id },
          files: ['./scripts/content.js'] // Asegúrate de que el nombre del archivo del content script sea correcto
        });
      });
    });
  } else {
    console.error('El botón de enviar campos no se encontró en el DOM.');
  }
});
