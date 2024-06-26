# Dockerfile

# Use the official Golang image as a build environment
FROM golang:1.22-alpine AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Use a minimal image for the final container
FROM alpine:3.20

# Install curl, ca-certificates, and envsubst (from gettext package)
RUN apk update && apk add --no-cache curl ca-certificates gettext

# Download and install the migrate tool
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz -o migrate.tar.gz \
    && tar -xzvf migrate.tar.gz -C /usr/local/bin migrate \
    && rm migrate.tar.gz

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=build /app/main .

# Copy the migrations folder
COPY migrations ./migrations

ARG DB_HOST=localhost
ARG DB_USER=root
ARG DB_PASSWORD=iloveu2much
ARG DB_NAME=todo

# Create a config template file
RUN echo "app:" > config.yaml \
    && echo "  name: Hugeman Exam" >> config.yaml \
    && echo "  port: 8080" >> config.yaml \
    && echo "" >> config.yaml \
    && echo "db:" >> config.yaml \
    && echo "  host: $DB_HOST" >> config.yaml \
    && echo "  port: 5432" >> config.yaml \
    && echo "  user: $DB_USER" >> config.yaml \
    && echo "  password: $DB_PASSWORD" >> config.yaml \
    && echo "  dbname: $DB_NAME" >> config.yaml

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["sh", "-c", "migrate -path ./migrations -database postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:5432/$DB_NAME?sslmode=disable up && ./main"]
