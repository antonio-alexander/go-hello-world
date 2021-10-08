# go-hello-world (github.com/antonio-alexander/go-hello-world)

This is meant to be a semi-simple hello-world program demonstration some basic go functionality, indicating the environment and methods for deployment/containerization. There are two applications, one for the cli (command line interface) and another for rest/http access. Both can be debugged, both as source code and as containers using the included dockerfiles and launch.json (vscode) to debug from the IDE.

## How to Use

To use, clone the repo into a folder, and the two commands below:

* docker-compose build
* docker-compose up -d

Once the container is up and running you can connect to it directly from VS code by using the "Attach (Remote)" configuration and/or view the output using the docker logs command.
