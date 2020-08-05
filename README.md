# Residentials

Sample project that helps with the management of all activities required for a
residential unit. This project parts is divided in several context, and this part
will be focused to Residential.

## Pre-requisites
You will need:

- [docker 19.03.12](https://docs.docker.com/get-docker/) or above.
- [docker-compose 1.25.4](https://docs.docker.com/compose/install/) or above.
- [Go 1.14](https://golang.org/doc/install) or above.

## Acknowledgment
This project was built using a [go-build-template](https://github.com/thockin/go-build-template)
that allows us to use a Makefile to drive the build (the universal API to software projects)
and a Dockerfile to build a docker image.

This has only been tested on Linux, and depends on Docker to build.

## Building

Run `make` or `make build` to compile your app.  This will use a Docker image
to build your app, with the current directory volume-mounted into place.  This
will store incremental state for the fastest possible build.  Run `make
all-build` to build for all architectures.

Run `make container` to build the container image.  It will calculate the image
tag based on the most recent git tag, and whether the repo is "dirty" since
that tag (see `make version`).  Run `make all-container` to build containers
for all supported architectures.

Run `make clean` to clean up.

Run `make help` to get a list of available targets.
