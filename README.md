# SSH Provision
[![Version](https://img.shields.io/github/v/release/tomdewildt/ssh-provision?label=version)](https://github.com/tomdewildt/ssh-provision/releases)
[![Build](https://img.shields.io/github/actions/workflow/status/tomdewildt/ssh-provision/ci.yml?branch=master)](https://github.com/tomdewildt/ssh-provision/actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/actions/workflow/status/tomdewildt/ssh-provision/cd.yml?label=release)](https://github.com/tomdewildt/ssh-provision/actions/workflows/cd.yml)
[![Coverage](https://img.shields.io/codecov/c/gh/tomdewildt/ssh-provision)](https://codecov.io/gh/tomdewildt/ssh-provision)
[![Report](https://goreportcard.com/badge/github.com/tomdewildt/ssh-provision)](https://goreportcard.com/report/github.com/tomdewildt/ssh-provision)
[![License](https://img.shields.io/github/license/tomdewildt/ssh-provision)](https://github.com/tomdewildt/ssh-provision/blob/master/LICENSE)

A simple tool to create ssh keys and distribute them to CentOS servers.

**Note:** This tool is not production ready.

# How To Run

Prerequisites:
* vagrant version ```2.2.7``` or later
* go version ```1.19``` or later

### Development

1. Run ```make init``` to initialize the environment.
2. Run ```make run``` to execute the cli tool.

### Test

1. Run ```make init``` to initialize the environment.
2. Run ```make test``` to execute the tests for the cli tool.

Run ```make vm/start``` to create the virtual machine, ```make vm/stop``` to stop the virtual machine and ```make vm/remove``` to remove it.

# How To Build

### Linux

1. Run ```make init``` to initialize the environment.
2. Run ```make build``` to build the cli tool.

# References

[Go Docs](https://golang.org/doc/)

[Go SSH](https://godoc.org/golang.org/x/crypto/ssh)

[Go Logus](https://godoc.org/github.com/sirupsen/logrus)

[Go Cobra](https://godoc.org/github.com/spf13/cobra)

[Go Validator](https://godoc.org/github.com/go-playground/validator)

[Go Testify](https://godoc.org/github.com/stretchr/testify)

[Go Project Layout](https://github.com/golang-standards/project-layout)
