#---------------------------------------------------------------------------------------------------
# hello-world-rest [Dockerfile]
#---------------------------------------------------------------------------------------------------

FROM golang:alpine AS build

WORKDIR /go/src/go-hello-world

COPY . /go/src/go-hello-world

RUN \
    cd cmd/rest \
    && go build -o hello-world-rest \
    && chmod +X hello-world-rest \
    && tar -czvf hello-world-rest.tar.gz ./hello-world-rest

FROM alpine:latest

WORKDIR /app

COPY --from=build /go/src/go-hello-world/cmd/rest/hello-world-rest.tar.gz /app/hello-world-rest.tar.gz

CMD tar -xzvf hello-world-rest.tar.gz && ./hello-world-rest
