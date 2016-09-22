ji-docker-2016
==============

`ji-docker-2016` holds the code and exercizes for
[docker@JI-2016](http://talks.godoc.org/github.com/sbinet/talks/2016/20160928-ji-docker/talk.slide)

The programming language used is [Go](https://golang.org).
You won't be programming in [Go](https://golang.org) though, just modifying some
strings being printed out.

## Installation & pre-requisites

You will need the `docker` container engine properly installed on your machine.
This is explained on the docker website:

- https://docs.docker.com/engine/installation

## Description

This repository holds 2 applications:

- `hello`, a command-line application that prints a greeting on the screen and tries to launch a sub-process (`pkg-config`, so nothing dangerous)
- `web-app`, a simple `http` server that serves a single page at `0.0.0.0:8080`

You *WON'T* need to install `Go` on your machine for this `ji-docker-2016` tutorial:
`Go` will be installed in the container.

The following are commands that should work *INSIDE* the container we will
build and deploy, and are thus *PURELY* *INFORMATIONAL*...

To build and install the `hello` command:

```sh
$> cd ji-docker-2016
$> go install ./hello
$> hello
hello JI-2016!
running pkg-config --version systemd now...
0.28
```

Similarly, for the `web-app` server, this time using the full repository name:

```sh
$> go install github.com/sbinet/ji-docker-2016/web-app
$> web-app
2016/09/20 16:43:58 listening on: http://localhost:8080
```

Another possibility is to use `go run` that compiles on the fly the executable and runs it:

```sh
$> go run ./web-app/main.go
2016/09/16 16:43:58 listening on: http://localhost:8080
```
