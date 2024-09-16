## Estructura del Proyecto

Aquí tienes una visión general de la estructura del directorio del proyecto:

- **`/cmd`**: El punto de entrada de la aplicación. Aquí se inicializan los servicios, el transporte y otros componentes.

- **`/pkg/service`**: Contiene la lógica de negocio de la aplicación, incluyendo la creación de usuarios, validaciones y otras funcionalidades centrales.

- **`/pkg/endpoint`**: Define los endpoints que exponen las funciones del servicio. Aquí es donde se procesa el JSON del frontend y se usa en la lógica de negocio.

- **`/pkg/transport`**: Contiene los controladores HTTP que gestionan las solicitudes de los clientes.

- **`/pkg/db`**: Gestiona las conexiones y operaciones de la base de datos.

- **`/models`**: Define los modelos de datos que representan las tablas de la base de datos.

- **`/config`**: Almacena las configuraciones del sistema, incluyendo credenciales de la base de datos y otras configuraciones.

## Uso del `Makefile`
### 1. Instalación

Para instalar `golangci-lint`, ejecuta el siguiente comando:

```sh
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### 2. **Asegúrate de que el BINARIO Está en tu PATH**

Después de instalar `golangci-lint`, el ejecutable se instalará en el directorio `$(go env GOPATH)/bin`. Asegúrate de que este directorio esté en tu PATH.

Para añadirlo a tu PATH, puedes modificar el archivo `.bashrc` (o `.zshrc`, si usas Zsh):

```sh
export PATH=$PATH:$(go env GOPATH)/bin
```

Después de hacer este cambio, ejecuta:

```sh
source ~/.bashrc
```

### 3. **Verifica la Instalación**

Para verificar si `golangci-lint` está en el PATH, ejecuta:

```sh
golangci-lint --version
```

Deberías ver la versión instalada. Si ves un error, puede ser que el PATH no se haya actualizado correctamente.

---
El `Makefile` se utiliza para automatizar tareas comunes en el desarrollo del proyecto. Incluye comandos para formatear el código, ejecutar análisis estático, y más. A continuación se explica cada objetivo del `Makefile`:

- **`fmt`**: Formatea el código fuente del proyecto usando `go fmt`, asegurando que el código esté limpio y siga el estilo de Go.

- **`vet`**: Ejecuta `go vet` para detectar posibles problemas en el código, como errores de formato o de uso del paquete estándar de Go.

- **`lint`**: Usa `golangci-lint`, una herramienta que combina múltiples linters para realizar un análisis más exhaustivo del código. Detecta problemas de estilo, errores potenciales, y prácticas de codificación que pueden mejorarse. `golangci-lint` incluye `golint`, `govet`, y otras herramientas útiles.

- Para ejecutar la aplicación y ejecutar análisis estático, usa:, usa:
  ```bash
  make run
  ```

---

## Configuración del Cloud SQL Proxy

Para conectar tu aplicación con una base de datos en Google Cloud SQL desde un sistema basado en Unix, sigue estos pasos para configurar y ejecutar el Cloud SQL Proxy:

1. **Descarga y configura el Cloud SQL Proxy**

   Ejecuta los siguientes comandos para descargar el Cloud SQL Proxy, hacer que el archivo sea ejecutable, y crear el directorio necesario:

   ```bash
   # Descarga el Cloud SQL Proxy
   wget https://dl.google.com/cloudsql/cloud_sql_proxy.linux.amd64 -O cloud_sql_proxy
   
   # Haz que el archivo sea ejecutable
   chmod +x cloud_sql_proxy
   
   # Crea el directorio para el socket
   sudo mkdir -p /cloudsql
   
   # Da permisos al directorio
   sudo chmod 777 /cloudsql
   ```

2. **Ejecuta el Cloud SQL Proxy**

   Una vez configurado, ejecuta el Cloud SQL Proxy usando el siguiente comando, reemplazando `cv-manager-432700:us-east1:cv-manager-db` con tu instancia de Cloud SQL:

   ```bash
   ./cloud_sql_proxy -dir=/cloudsql -instances=cv-manager-432700:us-east1:cv-manager-db
   ```

   Este comando iniciará el Cloud SQL Proxy y escuchará las conexiones en el directorio `/cloudsql`, que tu aplicación puede usar para conectarse a la base de datos en Google Cloud SQL.

---
**Nota:** El archivo `.env.example` es un ejemplo. Copia este archivo a `.env` y completa los valores con la información específica de tu entorno.