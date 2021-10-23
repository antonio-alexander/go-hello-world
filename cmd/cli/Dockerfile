#---------------------------------------------------------------------------------------------------
# hello-world-cli [Dockerfile]
# 
# Reference: https://stackoverflow.com/questions/63178036/how-to-find-commit-hash-from-within-a-running-docker-image
# commit: git rev-parse HEAD
# 
# https://stackoverflow.com/questions/6245570/how-to-get-the-current-branch-name-in-git
# branch: git rev-parse --abbrev-ref HEAD
#---------------------------------------------------------------------------------------------------

ARG GIT_BRANCH=no_git_branch_provided
ARG GIT_COMMIT=no_git_commit_provided

FROM golang:alpine AS build

ARG GIT_BRANCH
ARG GIT_COMMIT

WORKDIR /go/src/go-hello-world

COPY . /go/src/go-hello-world

RUN apk add jq 

RUN \
    VERSION=`cat /go/src/go-hello-world/version.json | jq '.Version' | sed 's/"//g'` \
    && cd cmd/cli \
    && go build -ldflags "-X github.com/antonio-alexander/go-hello-world/internal/cli.Version=${VERSION} -X github.com/antonio-alexander/go-hello-world/internal/cli.GitCommit=${GIT_COMMIT} -X github.com/antonio-alexander/go-hello-world/internal/cli.GitBranch=${GIT_BRANCH}" -o hello-world-cli \
    && chmod +X hello-world-cli \
    && tar -czvf hello-world-cli.tar.gz ./hello-world-cli

FROM alpine:latest

ARG GIT_BRANCH
ARG GIT_COMMIT

WORKDIR /app

COPY --from=build /go/src/go-hello-world/cmd/cli/hello-world-cli.tar.gz /app/hello-world-cli.tar.gz
COPY ./cmd/cli/sample_read_only.txt /app/sample_read_only.txt

LABEL antonio-alexander.git.branch=${GIT_BRANCH}
LABEL antonio-alexander.git.commit=${GIT_COMMIT}
LABEL org.opencontainers.image.source=https://github.com/antonio-alexander/go-hello-world

CMD tar -xzvf hello-world-cli.tar.gz && ./hello-world-cli
