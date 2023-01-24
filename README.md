# Go-Lambda

This repo is an example repo for deploying a golambda to AWS using AWS CDK.

Steps:

1. install the packages

```shell
xc get
```
or 

```shell
go get ./...
```

2. configure aws

```shell
aws configure
```

or set the following environment variables in your command line.

```shell
export AWS_ACCESS_KEY_ID="..."
export AWS_SECRET_ACCESS_KEY="..."
export AWS_SESSION_TOKEN="..."
```

You will need an IAM user to access these. After logging in, click a user, Then click `Command line or programmatic access`. Copy your environment variables.

3. run bootstrap

```shell
cdk bootstrap
```

4. deploy

```shell
cdk deploy
```




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