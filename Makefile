name := flax
build_path := ./build
coverage_path := ./coverage

docker_tag ?= latest
docker_image ?= moorara/$(name)


clean:
	@ rm -rf *.log $(name) $(build_path) $(coverage_path)

dep:
	@ dep ensure -update

run:
	@ go run cmd/main.go

build:
	@ ./scripts/build.sh --main ./cmd/main.go --binary ./$(name)

build-all:
	@ ./scripts/build.sh --all --main ./cmd/main.go --binary $(build_path)/$(name)

test:
	@ go test -v -race ./...

test-short:
	@ go test -race -short ./...

coverage:
	@ ./scripts/test-unit-cover.sh

docker:
	@ docker build -t $(docker_image):$(docker_tag) .

push:
	@ docker push $(docker_image):$(docker_tag)


.PHONY: clean
.PHONY: dep run build build-all
.PHONY: test test-short coverage
.PHONY: docker push
