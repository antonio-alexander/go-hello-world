#build from alpine container with golang
FROM golang:alpine AS build

#set the working directory to cmd, so we can build easily
WORKDIR /src/hello-http-rest/cmd

#set environment
ENV GO111MODULE auto

#copy current directory into the gopath
COPY . /src/hello-http-rest

#install delve and build the hell-http app
RUN \
    mkdir -p /app \
    go build -gcflags "all=-N -l"  -o /app/hello-world-rest \
    && chmod +x /app/hello-world-rest 

#build from a clean container (without golang)
FROM alpine:latest

#set working directory to app
WORKDIR /app

#copy the app
COPY --from=build /app/hello-http /app/hello-http

#use this to run delve with the executable
CMD /app/hello-world-rest

#use this to wait without doing anything if you want to interact with the container interactively
# docker exec -it hello-http /bin/ash
# CMD tail -f /dev/null