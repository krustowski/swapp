#
# swapp / Dockerfile
#

# https://hub.docker.com/_/golang

ARG GOLANG_VERSION
FROM golang:${GOLANG_VERSION}-alpine

ARG APP_NAME
ARG APP_FLAGS
ARG DOCKER_DEV_PORT
ARG TZ

ENV APP_NAME ${APP_NAME}
ENV APP_FLAGS ${APP_FLAGS}
ENV DOCKER_DEV_PORT ${DOCKER_DEV_PORT}
ENV TZ ${TZ}

RUN apk --no-cache add tzdata git

WORKDIR /go/src/${APP_NAME}
COPY . .

#RUN go mod init ${APP_NAME}
RUN go mod tidy
#RUN go install

# build the client -- wasm binary
RUN GOARCH=wasm GOOS=js go build -o web/app.wasm

# build the server
RUN go build ${APP_NAME}

EXPOSE ${DOCKER_DEV_PORT}
CMD ./${APP_NAME} ${APP_FLAGS}

