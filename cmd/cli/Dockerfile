#---------------------------------------------------------------------------------------------------
# hello-world-cli [Dockerfile]
#---------------------------------------------------------------------------------------------------

FROM golang:alpine AS build

WORKDIR /go/src/go-hello-world

COPY . /go/src/go-hello-world

RUN \
    cd cmd/cli \
    && go build -o hello-world-cli \
    && chmod +X hello-world-cli \
    && tar -czvf hello-world-cli.tar.gz ./hello-world-cli

FROM alpine:latest

WORKDIR /app

COPY --from=build /go/src/go-hello-world/cmd/cli/hello-world-cli.tar.gz /app/hello-world-cli.tar.gz
COPY ./sample_read_only.txt /app/sample_read_only.txt

CMD tar -xzvf hello-world-cli.tar.gz && ./hello-world-cli
