#---------------------------------------------------------------------------------------------------
# hello-world-cli [Dockerfile]
# 
# Reference: https://stackoverflow.com/questions/63178036/how-to-find-commit-hash-from-within-a-running-docker-image
# commit: git rev-parse HEAD
# 
# https://stackoverflow.com/questions/6245570/how-to-get-the-current-branch-name-in-git
# branch: git rev-parse --abbrev-ref HEAD
# 
# https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63
# The GOOS will always be linux, but in case you were curious
# 
# Sample docker build commands:
#  docker build -f ./cmd/cli/Dockerfile . -t ghcr.io/antonio-alexander/go-hello-world-cli:amd64_latest --build-arg GIT_COMMIT=$GITHUB_SHA --build-arg GIT_BRANCH=$GITHUB_REF --build-arg PLATFORM=linux/amd64 --build-arg GO_ARCH=amd64
#  docker build -f ./cmd/cli/Dockerfile . -t ghcr.io/antonio-alexander/go-hello-world-cli:armv7_latest --build-arg GIT_COMMIT=$GITHUB_SHA --build-arg GIT_BRANCH=$GITHUB_REF --build-arg PLATFORM=linux/arm/v7 --build-arg GO_ARCH=arm --build-arg GO_ARM=7
#
#---------------------------------------------------------------------------------------------------

ARG GIT_BRANCH=no_git_branch_provided
ARG GIT_COMMIT=no_git_commit_provided
ARG PLATFORM=linux/amd64
ARG GO_ARCH=amd64
ARG GO_ARM=7

FROM golang:alpine AS build

ARG GIT_BRANCH
ARG GIT_COMMIT
ARG GO_ARCH
ARG GO_ARM

WORKDIR /go/src/go-hello-world

COPY . /go/src/go-hello-world

RUN apk add jq

RUN \
    VERSION=`cat /go/src/go-hello-world/version.json | jq '.Version' | sed 's/"//g'` \
    && cd cmd/cli \
    && env GOARCH=${GO_ARCH} GOARM=${GO_ARM} GOOS=linux go build -ldflags "-X github.com/antonio-alexander/go-hello-world/internal.Version=${VERSION} -X github.com/antonio-alexander/go-hello-world/internal.GitCommit=${GIT_COMMIT} -X github.com/antonio-alexander/go-hello-world/internal.GitBranch=${GIT_BRANCH}" -o hello-world-cli \
    && chmod +X hello-world-cli \
    && tar -czvf hello-world-cli.tar.gz ./hello-world-cli

FROM --platform=${PLATFORM} alpine:latest

ARG GIT_BRANCH
ARG GIT_COMMIT

WORKDIR /app

COPY --from=build /go/src/go-hello-world/cmd/cli/hello-world-cli.tar.gz /app/hello-world-cli.tar.gz
COPY ./cmd/cli/sample_read_only.txt /app/sample_read_only.txt

LABEL antonio-alexander.git.branch=${GIT_BRANCH}
LABEL antonio-alexander.git.commit=${GIT_COMMIT}
LABEL org.opencontainers.image.source=https://github.com/antonio-alexander/go-hello-world

CMD tar -xzf hello-world-cli.tar.gz && ./hello-world-cli
