# Gin Boilerplate
> A skeleton to create an API with go and gin

[![Build Status][travis-image]][travis-url]
[![License: MIT][license-img]][license-url]

Golang gin boilerplate, with GORM, gin and tests !

<p align="center">
    <img alt="gopher logo" height="500px" src="logo.png">
</p>

## Quickstart

To run the server:

```bash
go run main.go
```

To run the tests:
```bash
mode=test go test -v ./...
```

## Configuration

The application load the `config.toml` if the environment variable mode is not defined.
If this variable is defined, the application will try to load the `config.[mode].toml` file.


[travis-image]: https://travis-ci.org/johnrazeur/gin-boilerplate.svg?branch=master
[travis-url]: https://travis-ci.org/johnrazeur/gin-boilerplate
[license-img]: https://img.shields.io/badge/License-MIT-yellow.svg
[license-url]: https://opensource.org/licenses/MIT
