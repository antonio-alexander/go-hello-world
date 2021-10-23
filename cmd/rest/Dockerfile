#---------------------------------------------------------------------------------------------------
# hello-world-rest [Dockerfile]
#---------------------------------------------------------------------------------------------------

ARG GIT_BRANCH=no_git_branch_provided
ARG GIT_COMMIT=no_git_commit_provided

FROM golang:alpine AS build

ARG GIT_BRANCH
ARG GIT_COMMIT

WORKDIR /go/src/go-hello-world

COPY . /go/src/go-hello-world

RUN \
    cd cmd/rest \
    && go build -ldflags "-X github.com/antonio-alexander/go-hello-world/internal/cli.Version=${VERSION} -X github.com/antonio-alexander/go-hello-world/internal/cli.GitCommit=${GIT_COMMIT} -X github.com/antonio-alexander/go-hello-world/internal/cli.GitBranch=${GIT_BRANCH}" -o hello-world-rest \
    && chmod +X hello-world-rest \
    && tar -czvf hello-world-rest.tar.gz ./hello-world-rest

FROM alpine:latest

WORKDIR /app

COPY --from=build /go/src/go-hello-world/cmd/rest/hello-world-rest.tar.gz /app/hello-world-rest.tar.gz

ARG GIT_BRANCH
ARG GIT_COMMIT

LABEL antonio-alexander.git.branch=${GIT_BRANCH}
LABEL antonio-alexander.git.commit=${GIT_COMMIT}
LABEL org.opencontainers.image.source=https://github.com/antonio-alexander/go-hello-world

CMD tar -xzvf hello-world-rest.tar.gz && ./hello-world-rest
