BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S_UTC')
GIT_COMMIT=$(shell git describe --tags --always --long --dirty)
LD_FLAGS=-X main.version=${GIT_COMMIT} -X main.buildTime=${BUILD_TIME}
GO_PKG=github.com/dailymotion/code-review-for-interviews
PACKAGES=$(shell go list ./... | grep -v '/vendor/')
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
GOVERSION=1.8
OUTPUT_DIR=output
OUTPUT_NAME=currencies
DOCKER_IMAGE_NAME=currencies
DOCKER_IMAGE_VERSION=${GIT_COMMIT}
DOCKER_REGISTRY=
DOCKER_TAG=${DOCKER_REGISTRY}${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_VERSION}

build:
	go get ${GO_PKG}/...
	GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags "${LD_FLAGS}" -o ${OUTPUT_DIR}/${OUTPUT_NAME} ${GO_PKG}

build-all:
	@$(MAKE) --no-print-directory build GOOS=linux   GOARCH=amd64 OUTPUT_NAME=${OUTPUT_NAME}-${GIT_COMMIT}-linux-amd64
	@$(MAKE) --no-print-directory build GOOS=darwin  GOARCH=amd64 OUTPUT_NAME=${OUTPUT_NAME}-${GIT_COMMIT}-darwin-amd64

build-all-in-docker:
	docker run --rm -v ${PWD}:/go/src/${GO_PKG} -w /go/src/${GO_PKG} golang:${GOVERSION} make build-all

build-in-docker:
	docker run --rm -v ${PWD}:/go/src/${GO_PKG} -w /go/src/${GO_PKG} golang:${GOVERSION} make build GOOS=${GOOS} GOARCH=${GOARCH}

build-static-linux-in-docker:
	docker run --rm -v ${PWD}:/go/src/${GO_PKG} -w /go/src/${GO_PKG} golang:${GOVERSION} make build GOOS=linux GOARCH=amd64 LD_FLAGS="${LD_FLAGS} -extldflags -static -linkmode external" OUTPUT_NAME=${OUTPUT_NAME}-${GIT_COMMIT}-static-linux-amd64

build-docker-image: build-static-linux-in-docker
	docker build --build-arg binary=${OUTPUT_DIR}/${OUTPUT_NAME}-${GIT_COMMIT}-static-linux-amd64 --tag ${DOCKER_TAG} .

clean:
	rm -rf ${OUTPUT_DIR}
