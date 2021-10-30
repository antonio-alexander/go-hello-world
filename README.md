# go-hello-world (github.com/antonio-alexander/go-hello-world)

go-hello-world contains two examples of Go applications that attempt to demo a cli (command line interface) REST application (web server). Each contains distribution files (Dockerfiles and docker-compose.yml) as well as a functional workflow.  This should serve as basic idea of how I should organize an environment around an application.

## Getting started

To get started with the code, clone the repository:

```sh
git clone git@github.com:antonio-alexander/go-hello-world.git
```

Once cloned, open the code in your IDE of choice and follow the instructions below as needed.

## Running as source code

To run as source code, navigate to the cmd folder, and within the cli or rest, you can run the code.

```sh
cd ./cmd/cli
go run main.go
```

This output is obviously redacted...because environmental variables can contain SECRETS...haha, even though they shouldn't.

```output
Hello, World!
Version: ""
Git Commit: ""
Git Branch: ""
 Present Working Directory: /Users/noobius/source_control/github.com/antonio-alexander/go-hello-world/cmd/cli
 Arguments: []
 Environmental Variables:
```

If you're using VS Code, there's some additional magic in the .vscode folder that can be used along with the launcher.

Keep in mind that the Version, Git Commit, and Git Branch are provided via ldflags on build, you'll need to build it with something like:

```sh
go build -ldflags "-X github.com/antonio-alexander/go-hello-world/internal.Version=${VERSION} -X github.com/antonio-alexander/go-hello-world/internal.GitCommit=${GIT_COMMIT} -X github.com/antonio-alexander/go-hello-world/internal.GitBranch=${GIT_BRANCH}" -o hello-world-cli
```

To get it to actually populate those values (I pulled this from the Dockerfile)

If you'd like to run the webserver as source code, you can do the following:

```sh
cd ./cmd/rest
go run main.go
```

```sh
starting web server on :8080
```

Once this executes, you can usse a web browser or the curl command to get feedback from the application:

```sh
curl http://localhost:8080
Hello, World!
Version: ""
Git Commit: ""
Git Branch: ""
```

Note that this will have the same issue as the cli when you build without the ldflags.

## Building images

Docker (or more accurately containers) are the distribution tool of choice for microservices. The docker images for the cli or rest can be build with either of the following commands.

This command builds ghcr.io/antonio-alexander/go-hello-world-cli:amd64_latest for intel

```sh
docker build -f ./cmd/cli/Dockerfile . -t ghcr.io/antonio-alexander/go-hello-world-cli:amd64_latest --build-arg GIT_COMMIT=$GITHUB_SHA --build-arg GIT_BRANCH=$GITHUB_REF --build-arg PLATFORM=linux/amd64 --build-arg GO_ARCH=amd64
```

This command builds ghcr.io/antonio-alexander/go-hello-world-cli:armv7_latest for the RPI4 (linux/armv7)

```sh
docker build -f ./cmd/cli/Dockerfile . -t ghcr.io/antonio-alexander/go-hello-world-cli:armv7_latest --build-arg GIT_COMMIT=$GITHUB_SHA --build-arg GIT_BRANCH=$GITHUB_REF --build-arg PLATFORM=linux/arm/v7 --build-arg GO_ARCH=arm --build-arg GO_ARM=7
```

The command above builds ghcr.io/antonio-alexander/go-hello-world-cli:amd64_latest for the RPI4 (linux/armv7)

```sh
docker build -f ./cmd/cli/Dockerfile . -t ghcr.io/antonio-alexander/go-hello-world-cli:armv7_latest --build-arg GIT_COMMIT=$GITHUB_SHA --build-arg GIT_BRANCH=$GITHUB_REF --build-arg PLATFORM=linux/arm/v7 --build-arg GO_ARCH=arm --build-arg GO_ARM=7
```

This command builds ghcr.io/antonio-alexander/go-hello-world-rest:amd64_latest for intel

```sh
docker build -f ./cmd/rest/Dockerfile . -t ghcr.io/antonio-alexander/go-hello-world-rest:amd64_latest --build-arg GIT_COMMIT=$GITHUB_SHA --build-arg GIT_BRANCH=$GITHUB_REF --build-arg PLATFORM=linux/amd64 --build-arg GO_ARCH=amd64
```

This command builds ghcr.io/antonio-alexander/go-hello-world-rest:armv7_latest for RPI4 (linux/armv7)

```sh
docker build -f ./cmd/rest/Dockerfile . -t ghcr.io/antonio-alexander/go-hello-world-rest:armv7_latest --build-arg GIT_COMMIT=$GITHUB_SHA --build-arg GIT_BRANCH=$GITHUB_REF --build-arg PLATFORM=linux/arm/v7 --build-arg GO_ARCH=arm --build-arg GO_ARM=7
```

These images can be built or pulled from ghcr.io

## Using the docker-compose

The docker-compose.yml can be used to build and run the images. Keep in mind that the latest tag references a manifest that will pull the appropriate platform (if present/supported) automatically. Keep in mind that by default it's configured for the intel image and if building for alternate platforms, you can comment/uncomment the image and argument lines.

To build:

```sh
docker compose -f ./cmd/docker-compose.yaml build
```

To pull:

```sh
docker compose -f ./cmd/docker-compose.yaml pull
```

To run (and not attach):

```sh
docker compose -f ./cmd/docker-compose.yaml up -d
```
