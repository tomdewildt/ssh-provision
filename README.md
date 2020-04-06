# SSH Provision
[![Go Version](https://img.shields.io/badge/go-v1.13-blue)](https://golang.org/dl/#stable)
![Project Version](https://img.shields.io/github/v/release/tomdewildt/ssh-provision?label=version)
![CI Status](https://github.com/tomdewildt/ssh-provision/workflows/ci/badge.svg?branch=master)
![CD Status](https://github.com/tomdewildt/ssh-provision/workflows/cd/badge.svg)
[![Coverage Status](https://codecov.io/gh/tomdewildt/ssh-provision/branch/master/graph/badge.svg)](https://codecov.io/gh/tomdewildt/ssh-provision)
[![Report Status](https://goreportcard.com/badge/github.com/tomdewildt/ssh-provision)](https://goreportcard.com/report/github.com/tomdewildt/ssh-provision)

A simple tool to create ssh keys and distribute them to CentOS servers.

**Note:** This tool is not production ready.

# How To Run

Prerequisites:
* vagrant version ```2.2.7``` or later
* go version ```1.13``` or later

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
