# Build Stage
FROM golang:1.26-alpine AS build-env

# Set environment variables
ENV APP_NAME=dnscheck
ENV CMD_PATH=main.go

# Maintainer Info
LABEL maintainer="Chiman Jain <chimanjain15@gmail.com>"

# Set working directory
WORKDIR /app

# Copy dependency files and download modules
COPY go.mod go.sum ./
RUN go mod download

# Copy application source
COPY . .

# Build application
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /${APP_NAME} ${CMD_PATH}

# Run Stage
FROM alpine

# Set environment variables
ENV APP_NAME=dnscheck

# Set working directory
WORKDIR /app

# Copy only the compiled binary from the build stage
COPY --from=build-env /${APP_NAME} .

# Expose application port
EXPOSE 3000

# Start app
CMD ["./dnscheck"]
