#
# swapp / Dockerfile
#

#
# stage 0 -- build
#

# https://hub.docker.com/_/golang
ARG GOLANG_VERSION
FROM golang:${GOLANG_VERSION}-alpine AS swapp-build

ARG APP_NAME

ENV APP_NAME ${APP_NAME}

RUN apk add git

WORKDIR /go/src/${APP_NAME}
COPY . .

#RUN go mod init ${APP_NAME}
RUN go mod tidy
#RUN go install

# build the client -- wasm binary
RUN GOARCH=wasm GOOS=js go build -o web/app.wasm -tags wasm

# build the server
RUN go build ${APP_NAME}


#
# stage 1 -- release
#

FROM alpine:3.16 AS swapp-release

ARG APP_FLAGS
ARG DOCKER_DEV_PORT

ENV APP_FLAGS ${APP_FLAGS}
ENV DOCKER_DEV_PORT ${DOCKER_DEV_PORT}

COPY web/ /opt/web/
COPY --from=swapp-build /go/src/swapp/swapp /opt/swapp
COPY --from=swapp-build /go/src/swapp/web/app.wasm /opt/web/app.wasm

WORKDIR /opt
EXPOSE ${DOCKER_DEV_PORT}
CMD ./swapp


