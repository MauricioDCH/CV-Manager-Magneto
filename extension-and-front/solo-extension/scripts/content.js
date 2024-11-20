let inputs = document.querySelectorAll('input[type="text"]');
let attributes = [];
let nameAttributes = ['aria-label', 'name', 'placeholder'];
let desiredAttributes = ['aria-label', 'name', 'placeholder', 'type'];

let filteredInputs = Array.from(inputs).filter(input => {
    return nameAttributes.some(attr => input.hasAttribute(attr));
});

const token = localStorage.getItem('token');

if (token) {
    console.log("Token encontrado en localStorage:", token);
    chrome.storage.local.set({ token: token }, () => {
        console.log('Token guardado en chrome.storage.local automáticamente');
    });
} else {
    console.error('No se encontró el token en localStorage');
}

// Itera sobre los elementos seleccionados y extrae los atributos deseados
filteredInputs.forEach(input => {
    let attr = {};
    Array.from(input.attributes).forEach(attribute => {
        if (desiredAttributes.includes(attribute.name)) {
            attr[attribute.name] = attribute.value;
        }
    });
    if (Object.keys(attr).length > 0) {
        attributes.push(attr);
    }
});

// Escucha los mensajes desde la página web
window.addEventListener('message', (event) => {
    if (event.source !== window) return;

    if (event.data.type) {
        if (event.data.type === 'FROM_PAGE') {
            let userInfo = event.data.userInfo;
            console.log('Datos del usuario recibidos:', userInfo);

            // Guarda el email en almacenamiento local
            chrome.storage.local.set({ email: userInfo.email }, () => {
                console.log('Email guardado:', userInfo.email);
            });
        }

        // Maneja el mensaje con tipo 'FILL_FORM'
        if (event.data.type === 'FILL_FORM') {
            chrome.storage.local.get('selectedCvId', (result) => {
                let selectedCvId = parseInt(result.selectedCvId) || 0;
                let inputsData = attributes;

                // Formatea los datos en el formato solicitado
                let requestData = {
                    "inputs": inputsData,
                    "idcv": selectedCvId
                };

                console.log('Atributos extraídos:', attributes);
                console.log('Datos a enviar:', requestData);

                // Enviar los datos al endpoint y recibir respuesta para rellenar el formulario
                //fetch('http://localhost:5000/endpoint', {
                fetch('http://34.45.83.31:5000/endpoint', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(requestData) // Enviar los datos formateados
                })
                .then(response => response.json())
                .then(data => {
                    console.log('Datos del servidor recibidos:', data);
                    fillFormFields(data);  // Rellenar el formulario con los datos recibidos
                })
                .catch(error => console.error('Error en la solicitud:', error));
            });
        }
    }
});

// Función para rellenar los campos con los datos del JSON
function fillFormFields(data) {
    data.inputs.forEach(inputData => {
        let input = document.querySelector(`input[name="${inputData.name}"], input[aria-label="${inputData['aria-label']}"]`);
        if (input) {
            input.value = inputData.value;
        }
    });
}
