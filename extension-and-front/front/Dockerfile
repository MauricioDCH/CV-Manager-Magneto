# Dockerfile para servicio del forn-end de la aplicación.

FROM node:18-alpine

WORKDIR /app

COPY . .

RUN npm install
RUN npm install --save-dev @babel/plugin-proposal-private-property-in-object
RUN npm install react-router-dom
RUN npm install jwt-decode
RUN npm install --save @fortawesome/fontawesome-free
RUN npm run build

EXPOSE 5173

ENTRYPOINT ["npm", "run", "dev", "--", "--host"]
