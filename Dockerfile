# Use an official Go runtime as the base image
FROM golang:1.19 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download and cache Go modules
RUN go mod download

# Copy the rest of the application code
COPY . .

RUN cd utils && ls -la

# Generate env variable for db credentiels we going to add default value just in case

ENV DB_USER=root

ENV DB_PASS=root

ENV DB_HOST=localhost

ENV DB_NAME=booky

#defaul port PG
ENV DB_PORT=5432

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /booky

RUN ls -la

#Expose teh default port of the app 

EXPOSE 8000

# Run the application
CMD ["/booky"]
