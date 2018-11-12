[![Build Status][travisci-image]][travisci-url]

# Flax

This is a **work-in-progress** utility for mocking APIs.

## TO-DO

Supporting the following API paradigms:

  - [ ] REST

Supporting the following features:

  - [ ] HTTPS

## Commands

| Command                        | Description                                          |
|--------------------------------|------------------------------------------------------|
| `make dep`                     | Install and updates dependencies                     |
| `make run`                     | Run the application locally                          |
| `make build`                   | Build the binary locally                             |
| `make build-all`               | Build the binary locally for all supported platforms |
| `make test`                    | Run the unit tests                                   |
| `make test-short`              | Run the unit tests using `-short` flag               |
| `make coverage`                | Run the unit tests with coverage report              |
| `make docker`                  | Build Docker image                                   |
| `make push`                    | Push built image to registry                         |


[travisci-url]: https://travis-ci.org/moorara/flax
[travisci-image]: https://travis-ci.org/moorara/flax.svg?branch=master
