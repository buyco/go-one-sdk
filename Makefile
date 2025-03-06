DOCKER_BUILD := docker run --rm -u `id -u` -v ${PWD}:/sdk openapitools/openapi-generator-cli:v7.12.0 generate -i sdk/api_files/one_track_and_trace.yaml
GO_CLIENT := -g go -o /sdk/one \
			--git-repo-id=go-one-sdk \
			--git-user-id=buyco \
			--additional-properties=packageName=one \
			--additional-properties=isGoSubmodule=true \
			--additional-properties=generateInterfaces=true
## generate: Clean and generate SDK from file.
generate:
	${MAKE} clean
	${MAKE} go-sdk
	${MAKE} fmt

go-sdk:
	${DOCKER_BUILD} ${GO_CLIENT}

clean:
	rm -rf ./one

fmt:
	gofmt -w -s ./one 1>&2

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
