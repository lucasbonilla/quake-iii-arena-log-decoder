# Build Stage
# First pull Golang image
FROM golang:1.22.3-bullseye as build
 
# Set environment variable
ENV APP_NAME quake-iii-arena-log-decoder
ENV CMD_PATH cmd/app/main.go
 
# Copy application data into image
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

# Budild application
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH

# Run Stage
FROM alpine:latest AS script
 
# Set environment variable
ENV APP_NAME quake-iii-arena-log-decoder
 
# Copy only required data into this image
COPY --from=build /$APP_NAME .
COPY ./config.toml .
RUN mkdir -p files/in/
COPY ./files/in/qgames.log ./files/in
 
# Start app
CMD ./$APP_NAME