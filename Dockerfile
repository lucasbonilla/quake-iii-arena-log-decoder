# ---- build ----
FROM golang:1.22.3-bullseye as build
ENV APP_NAME quake-iii-arena-log-decoder
ENV CMD_PATH cmd/app/main.go
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH

# ---- script ----
FROM alpine:latest AS script
ENV APP_NAME quake-iii-arena-log-decoder
COPY --from=build /$APP_NAME .
COPY ./config.toml .
RUN mkdir -p files/in/
COPY ./files/in/qgames.log ./files/in
CMD ./$APP_NAME