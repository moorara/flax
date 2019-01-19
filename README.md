[![Build Status][circleci-image]][circleci-url]
[![Go Report Card][goreport-image]][goreport-url]

# Flax

This is a **WORK-IN-PROGRESS**.

Flax is a service for mocking your APIs for testing purposes.

## Quick Start

### Docker

### Examples

## TO-DO

Supporting the following features:

  - **Connection**
    - [ ] HTTP
    - [ ] HTTPS
  - **Mocking**
    - [ ] Basic HTTP
    - [ ] RESTful HTTP
    - [ ] GraphQL
  - **Verifying**
    - [ ] *TBD*
  - **Configuration**
    - [x] YAML
    - [x] JSON
    - [ ] REST API

## Development

| Command            | Description                                          |
|--------------------|------------------------------------------------------|
| `make run`         | Run the application locally                          |
| `make build`       | Build the binary locally                             |
| `make build-all`   | Build the binary locally for all supported platforms |
| `make test`        | Run the unit tests                                   |
| `make test-short`  | Run the unit tests using `-short` flag               |
| `make coverage`    | Run the unit tests with coverage report              |
| `make docker`      | Build Docker image                                   |
| `make push`        | Push built image to registry                         |
| `make save-docker` | Save built image to disk                             |
| `make load-docker` | Load saved image from disk                           |


[circleci-url]: https://circleci.com/gh/moorara/flax/tree/master
[circleci-image]: https://circleci.com/gh/moorara/flax/tree/master.svg?style=shield

[goreport-url]: https://goreportcard.com/report/github.com/moorara/flax
[goreport-image]: https://goreportcard.com/badge/github.com/moorara/flax
