# Build Stage
# First pull Golang image
FROM golang:1.24-alpine as build-env

# Set envirment variable
ENV APP_NAME dnscheck
ENV CMD_PATH main.go

# Maintainer Info
LABEL maintainer="Chiman Jain <chimanjain15@gmail.com>"

# Copy application data into image
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

# Build application
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH

# Run Stage
FROM alpine

# Set envirment variable
ENV APP_NAME dnscheck

# Copy only required data into this image
COPY --from=build-env /$APP_NAME .

# Expose application port
EXPOSE 3000

# Start app
CMD ./$APP_NAME
