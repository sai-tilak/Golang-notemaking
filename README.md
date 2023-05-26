# Golang-notemaking
Using Go language created REST API's for end users 

Step 1: Create a Dockerfile
Create a file named "Dockerfile" in the same directory as your Go code with the following content:
# Start from a base Golang image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the remaining source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the desired port (change it according to your application)
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

Step 2: Build the Docker Image
Open a terminal or command prompt, navigate to the directory containing the Dockerfile, and run the following command:
docker build -t note-app .
This command builds a Docker image named "note-app" based on the Dockerfile.

Step 3: Run the Docker Container Locally
Once the Docker image is built, you can run a container locally using the following command:
docker run -p 8080:8080 note-app
This command starts a container from the "note-app" image and maps port 8080 from the container to port 8080 on your local machine. Adjust the port mapping as needed based on your application's configuration.

Step 4: Push the Docker Image to DockerHub
To host the Docker image on DockerHub, you need to create a repository on DockerHub and tag your local image accordingly. Follow these steps:

1. Log in to DockerHub using the Docker CLI:
docker login
2. Tag the local image with your DockerHub username and repository name:
docker tag note-app your-dockerhub-username/note-app:latest
Replace "your-dockerhub-username" with your actual DockerHub username.

3. Push the tagged image to DockerHub:
docker push your-dockerhub-username/note-app:latest

This command uploads the image to your DockerHub repository.

After successfully pushing the image to DockerHub, you can access it from anywhere.

That's it! Your Go application is now containerized using Docker, and the container image is hosted on DockerHub.






