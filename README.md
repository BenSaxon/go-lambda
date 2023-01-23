## Tasks

These tasks follow [eXeCute](https://github.com/Joe-Davidson1802/xc) syntax, therefore can be ran with `xc [taskname]`.

### get

Install the Go dependencies for the project.

```shell
go get ./...
```

### build-api

Builds the `api/`.

```shell
go build ./...
```

### run-api

Runs the api locally.

Requires: build-api

```shell
go run api/main.go
```