# Build Stage
# First pull Golang image
FROM golang:1.18 as build-env

# Set envirment variable
ENV APP_NAME dnscheck
ENV CMD_PATH main.go

# Copy application data into image
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

# Build application
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH

# Run Stage
FROM alpine:3.15

# Set envirment variable
ENV APP_NAME dnscheck

# Copy only required data into this image
COPY --from=build-env /$APP_NAME .

# Expose application port
EXPOSE 3000

# Start app
CMD ./$APP_NAME