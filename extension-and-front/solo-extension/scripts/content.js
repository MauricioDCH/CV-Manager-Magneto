//document.addEventListener('DOMContentLoaded', () => {
    // Selecciona todos los elementos de tipo input en la página
    let inputs = document.querySelectorAll('input[type="text"]');
    let attributes = [];
    let nameAttributes = ['aria-label', 'name', 'placeholder'];
    let desiredAttributes = ['aria-label', 'name', 'placeholder', 'type'];

    // Filtra los inputs para aquellos que tengan al menos uno de los atributos deseados
    let filteredInputs = Array.from(inputs).filter(input => {
        // Verifica si al menos uno de los atributos deseados está presente en el input
        return nameAttributes.some(attr => input.hasAttribute(attr));
    });

    // JSON quemado con los datos simulados
    let jsonData = {
        "inputs": [
            { "name": "fbclc_userName", "value": "somema" },
            { "name": "fbclc_emailConf", "value": "sofia@example.com" },
            { "name": "fbclc_fName", "value": "Sofia" },
            { "name": "fbclc_lName", "value": "Mendieta" },
            { "name": "fbclc_phoneNumber", "value": "3002008000" }
        ]
    };

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

    // Función para rellenar los campos con los datos del JSON
    function fillFormFields(data) {
        data.inputs.forEach(inputData => {
            let input = document.querySelector(`input[name="${inputData.name}"], input[aria-label="${inputData['aria-label']}"]`);
            if (input) {
                input.value = inputData.value;
            }
        });
    }

    // Escucha los mensajes desde la página web
    window.addEventListener('message', (event) => {
        // Verifica que el mensaje proviene de la misma página
        if (event.source !== window) return;

        // Verifica que el tipo de mensaje es correcto
        if (event.data.type && event.data.type === 'FROM_PAGE') {
            let userInfo = event.data.userInfo;
            console.log('Datos del usuario recibidos:', userInfo);

            // Aquí puedes almacenar el email para usarlo más tarde
            chrome.storage.local.set({ email: userInfo.email }, () => {
                console.log('Email guardado:', userInfo.email);
            });
        }

        // Escucha el mensaje para rellenar el formulario
        if (event.data.type && event.data.type === 'FILL_FORM') {
            fillFormFields(jsonData);  // Usar el JSON quemado
        }

        chrome.storage.local.get('email', (result) => {
            let emailValue = result.email || ''; // Si no hay email, usa una cadena vacía
            //let inputsData = JSON.stringify(attributes, null, 2);
            let inputsData = attributes;
            // Formatea los datos en el formato solicitado
            let requestData = {
                "inputs": inputsData,     // Lista de inputs con sus atributos
                "email": emailValue       // Valor del campo email
            };

            console.log('Atributos extraídos:', attributes);
            console.log('Datos a enviar:', requestData);

        /*
            // Agrega un campo de email (esto se puede ajustar según la lógica de tu aplicación)
            let emailField = document.querySelector('input[name="email"]');
            let emailValue = emailField ? emailField.value : '';
            let inputsData = attributes;

            // Formatea los datos en el formato solicitado
            let requestData = {
                "inputs": inputsData,     // Lista de inputs con sus atributos
                "email": emailValue       // Valor del campo email
            };

            console.log('Atributos extraídos:', attributes);

            console.log('Datos a enviar:', requestData);
        */
            // Enviar datos al endpoint
        /* fetch('http://localhost:5000/endpoint', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(requestData)
            })
            .then(response => {
                console.log('Respuesta del servidor:', response);
                return response.json();
            })
            .then(data => console.log('Datos del servidor:', data))
            .catch(error => console.error('Error en la solicitud:', error));*/
    });
});