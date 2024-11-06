# CV-Manager-Magneto
## 1. Authors

| [<img src="https://avatars.githubusercontent.com/u/89171062?v=4" width=115><br><sub>Brayan David Zuluaga Giraldo.</sub>](https://github.com/bdzuluagag) |  [<img src="https://avatars.githubusercontent.com/u/79530549?v=4" width=115><br><sub>Sofía Mendieta Marín</sub>](https://github.com/somendietam) |   [<img src="https://avatars.githubusercontent.com/u/81777898?s=400&u=2eeba9c363f9c474c7fb419ef36562e2d2b6b866&v=4" width=115><br><sub>Mauricio David Correa Hernández.</sub>](https://github.com/MauricioDCH) |   [<img src="https://avatars.githubusercontent.com/u/56942218?v=4" width=115><br><sub>Pascual Gómez Londoño.</sub>](https://github.com/pascualgomz) |   [<img src="https://avatars.githubusercontent.com/u/100231247?v=4" width=115><br><sub>Juan Manuel Lopez Sanchez.</sub>](https://github.com/JuanMaLopez2) | 
| :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------: |:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------: |:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------: |:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------: |:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------: |

# ENDPOINTS AND HOW TO RUN IT.

## WITH DOCKER.

### **Use the command below**

Make sure you have docker installed on your machine.

[Install Docker](https://www.docker.com/)

```bash
docker compose up --build
```

## WITHOUT DOCKER.

For more information, go to each service.

### Open 7 terminals, then run all services below.

### **BEFORE RUN SERVICES, RUN THE CLOUD SQL PROXY**

You only need to run once the proxy.


```bash
./cloud_sql_proxy -dir=/cloudsql -instances=cv-manager-432700:us-east1:cv-manager-db
```
---
### CREATE CV.

**Endpoint**
```bash
http://localhost:8081/create-cv
```

**To run**
```bash
npm run dev
```

---

### CV.

**Endpoint**
```bash
http://localhost:8008/cv/{id}
http://localhost:8008/cv/user/{user_id}
```
**To run**
```bash
npm run dev
```

---

### EXTENSION AND FRONT.
### Extension.

**Endpoint**
```bash
http://localhost:5000/endpoint
```

**To run**

- Install the extension in the extensions settings.
    ```bash
    chrome://extensions/
    ```

- Localize `Load unpacked`.

- Serch for in the namespace the folder named `solo-extension file`.

- Run it in the extension location of your browser.


### Front.

**Endpoint**
```bash
http://localhost:5173/
```

**To run**
```bash
npm run dev
```

---

### LOGIN.
**Endpoint**
```bash
http://localhost:8000/login
```

**To run**
```bash
make run
```

---

### REGISTER.
**Endpoint**
```bash
http://localhost:8080/register
```

**To run**
```bash
make run
```

---

### SERVER.
**Endpoint**
```bash
http://localhost:5000/endpoint
```

**To run**
```bash
make run
```

---