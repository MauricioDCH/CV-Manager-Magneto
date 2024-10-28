# Use an official Node.js runtime as a parent image
FROM node:18-alpine

# Set the working directory inside the container
WORKDIR /app

COPY . .

# Install dependencies
RUN npm install
RUN npm install --save-dev @babel/plugin-proposal-private-property-in-object


# Build the React app
RUN npm run build

# Make port 5173 available to the world outside this container
EXPOSE 5173

# Command to run the application
ENTRYPOINT ["npm", "run", "dev", "--", "--host"]
