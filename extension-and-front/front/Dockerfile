# Dockerfile para servicio del forn-end de la aplicación.

FROM node:18-alpine

WORKDIR /app

COPY . .

RUN npm install
RUN npm install --save-dev @babel/plugin-proposal-private-property-in-object

RUN npm run build

EXPOSE 5173

ENTRYPOINT ["npm", "run", "dev", "--", "--host"]
