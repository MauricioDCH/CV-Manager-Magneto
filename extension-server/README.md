# CV-Manager-Magneto
## 1. Authors
---
| [<img src="https://avatars.githubusercontent.com/u/89171062?v=4" width=115><br><sub>Brayan David Zuluaga Giraldo.</sub>](https://github.com/bdzuluagag) |  [<img src="https://avatars.githubusercontent.com/u/79530549?v=4" width=115><br><sub>Sofía Mendieta Marín</sub>](https://github.com/somendietam) |   [<img src="https://avatars.githubusercontent.com/u/81777898?s=400&u=2eeba9c363f9c474c7fb419ef36562e2d2b6b866&v=4" width=115><br><sub>Mauricio David Correa Hernández.</sub>](https://github.com/MauricioDCH) |   [<img src="https://avatars.githubusercontent.com/u/56942218?v=4" width=115><br><sub>Pascual Gómez Londoño.</sub>](https://github.com/pascualgomz) |   [<img src="https://avatars.githubusercontent.com/u/100231247?v=4" width=115><br><sub>Juan Manuel Lopez Sanchez.</sub>](https://github.com/JuanMaLopez2) | 
| :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------: |:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------: |:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------: |:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------: |:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------: |

## SERVER FOR EXTENSION
This is the server for extension queries to the Database and Gemini.

### Dependencies
In order to run this project, you will need to execute the commands to obtain the required dependencies.


```bash
go get cloud.google.com/go/cloudsqlconn
go get github.com/jackc/pgx/v5
go get github.com/jackc/pgx/v5/stdlib
go get github.com/google/generative-ai-go/genai
go get google.golang.org/api/option
go get github.com/joho/godotenv
go get gorm.io/driver/postgres
```

To ensure there are no errors with dependencies, please execute:
```bash
go mod tidy
```

---

### Setting up the Cloud SQL Proxy

To connect your application with a database in Google Cloud SQL from a Unix-based system, follow these steps to configure and run the Cloud SQL Proxy:

1. **Download and configure the Cloud SQL Proxy**

   Run the following commands to download the Cloud SQL Proxy, make the file executable, and create the necessary directory:

   ```bash
    # Download the Cloud SQL Proxy
    wget https://dl.google.com/cloudsql/cloud_sql_proxy.linux.amd64 -O cloud_sql_proxy

    # Make the file executable
    chmod +x cloud_sql_proxy

    # Create the directory for the socket
    sudo mkdir -p /cloudsql

    # Give permissions to the directory
    sudo chmod 777 /cloudsql
   ```

2. **Run the Cloud SQL Proxy**

   Once configured, run the Cloud SQL Proxy using the following command, replacing `cv-manager-432700:us-east1:cv-manager-db` with your Cloud SQL instance:

   ```bash
   ./cloud_sql_proxy -dir=/cloudsql -instances=cv-manager-432700:us-east1:cv-manager-db
   ```

    This command will start the Cloud SQL Proxy and listen for connections in the `/cloudsql` directory, which your application can use to connect to the database in Google Cloud SQL.

### To run the server.


Before running the server make sure the database is working.

If the file named .env doesn't exist, change the .env_example file name to .env and introduce the credentials.

Make sure the makefile is in the main folder and contains the following content.

```makefile
.PHONY: fmt vet lint run

fmt:
	go fmt ./...

vet:
	go vet ./...

run: fmt vet
	go run cmd/main.go
```


Then, execute the command below.
```bash
make run
```

The server listens on port 5000.

If you whant to send some POST, you can execute in POSTMAN.

```bash
http://localhost:5000/endpoint
```

### Request example.
In the body of the post, use this example.
```json
{
    "inputs": [
        {
            "name": "name",
            "type": "text"
        },
        {
            "name": "last_name",
            "type": "text"
        },
        {
            "name": "email",
            "type": "text"
        },
        {
            "name": "education",
            "type": "text"
        },
        {
            "name": "skills",
            "type": "text"
        },
        {
            "name": "pets",
            "type": "text"
        },
        {
            "name": "psiquiatric-issues",
            "type": "text"
        },
        {
            "name": "additional-information-for-postulation",
            "type": "text"
        }
    ],
    "email": "sergio.munoz@hotmail.com"
}
```

### This is an response example obtained
```json
{
    "inputs": [
        {
            "name": "name",
            "value": "Juan"
        },
        {
            "name": "last_name",
            "value": "Pérez"
        },
        {
            "name": "email",
            "value": "juan.perez@example.com"
        },
        {
            "name": "education",
            "value": "Ingeniería de Sistemas"
        },
        {
            "name": "skills",
            "value": "Go, Python, SQL"
        },
        {
            "name": "pets",
            "value": "No tengo mascotas."
        },
        {
            "name": "psiquiatric-issues",
            "value": "No tengo ningún problema de salud mental."
        },
        {
            "name": "additional-information-for-postulation",
            "value": "Estoy buscando un trabajo desafiante que me permita poner en práctica mis habilidades y crecer profesionalmente."
        }
    ]
}
```