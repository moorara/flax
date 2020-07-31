[![Build Status][workflow-image]][workflow-url]
[![Go Report Card][goreport-image]][goreport-url]
[![Test Coverage][coverage-image]][coverage-url]
[![Maintainability][maintainability-image]][maintainability-url]

# Flax

This is a **WORK-IN-PROGRESS**.

Flax is a service for mocking your APIs for testing purposes.

## Quick Start

### Docker

Create a _spec_ file named `flax.yaml` as follows:

```yaml
http:
  - path: /health
  - methods: [ POST ]
    path: /api/v1/sendMessage
    response:
      status: 201
      body: {
        "id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"
      }

rest:
  - base_path: /api/v1/teams
    store:
      objects: [
        { "_id": "aaaaaaaaaaaaaaaaaaaaaaaa", "name": "Back-end" },
        { "_id": "bbbbbbbbbbbbbbbbbbbbbbbb", "name": "Front-end" }
      ]
```

Then run the following command:

```
docker container run -d -p "8080:8080" -p "9999:9999" -v "$PWD/flax.yaml:/flax.yaml" moorara/flax:latest
```

Now, open your browser and hit `http://localhost:8080/api/v1/teams`.

### Examples

You can find more examples [here](./examples).

## TO-DO

Supporting the following features:

  - **Connection**
    - [x] HTTP
    - [ ] HTTPS (TLS, mTLS)
  - **Mocking**
    - [x] Basic HTTP
    - [x] RESTful HTTP
  - **Configuration**
    - [x] YAML Spec
    - [x] JSON Spec
    - [ ] REST API
  - **Verification**
    - [ ] REST API

## Development

| Command            | Description                                          |
|--------------------|------------------------------------------------------|
| `make build`       | Build the binary locally                             |
| `make build-all`   | Build the binary locally for all supported platforms |
| `make test`        | Run the unit tests                                   |
| `make test-short`  | Run the unit tests using `-short` flag               |
| `make coverage`    | Run the unit tests with coverage report              |
| `make docker`      | Build Docker image                                   |
| `make push`        | Push built image to registry                         |
| `make save-docker` | Save built image to disk                             |
| `make load-docker` | Load saved image from disk                           |


[workflow-url]: https://github.com/moorara/flax/actions
[workflow-image]: https://github.com/moorara/flax/workflows/Master/badge.svg
[goreport-url]: https://goreportcard.com/report/github.com/moorara/flax
[goreport-image]: https://goreportcard.com/badge/github.com/moorara/flax
[coverage-url]: https://codeclimate.com/github/moorara/flax/test_coverage
[coverage-image]: https://api.codeclimate.com/v1/badges/3c6a95f727fc89be77eb/test_coverage
[maintainability-url]: https://codeclimate.com/github/moorara/flax/maintainability
[maintainability-image]: https://api.codeclimate.com/v1/badges/3c6a95f727fc89be77eb/maintainability
