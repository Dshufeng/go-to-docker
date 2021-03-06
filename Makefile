# Inspired from https://github.com/philips/grpc-gateway-example

GOPATHP?=/Users/fanhongling/Downloads/workspace
GOPATHD?=/home/vagrant/go

IMG_NS?=docker.io/tangfeixiong
IMG_REPO?=go-to-docker
IMG_TAG?=0.1
GIT_COMMIT=$(shell date +%y%m%d%H%M)-git$(shell git rev-parse --short=7 HEAD)
REGISTRY_HOST?=172.17.4.50:5000

all: protoc-grpc docker-push

protoc-grpc: protoc-moby
	protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gofast_out=Mpb/moby/api.proto=github.com/tangfeixiong/go-to-docker/pb/moby,Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
		pb/service.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=Mpb/moby/api.proto=github.com/tangfeixiong/go-to-docker/pb/moby,logtostderr=true:. \
		pb/service.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:. \
		pb/service.proto
	go generate ./pb

protoc-moby:
	protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gofast_out=Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,Mpb/moby/api.proto=github.com/tangfeixiong/go-to-docker/pb/moby:. \
        pb/moby/api.proto

go-install:
	go install -v ./

go-build:
	@CGO_ENABLED=0 go build -a -v -o ./bin/gotodocker --installsuffix cgo -ldflags "-s" ./
	#@CGO_ENABLED=0 go build -v -o ./bin/gotodocker -tags netgo -installsuffix netgo -ldflags "-s" ./

docker-build: go-build
	docker build -t $(IMG_NS)/$(IMG_REPO):$(IMG_TAG) ./

docker-push: docker-build
	docker push $(IMG_NS)/$(IMG_REPO):$(IMG_TAG)

docker-cgo:
	go build -v -a -o ./bin/gotodocker ./
	docker build -t $(IMG_NS)/$(IMG_REPO):$(IMG_TAG)-$(GIT_COMMIT) -f Dockerfile.cgo ./
	docker push $(IMG_NS)/$(IMG_REPO):$(IMG_TAG)

docker-run:
	$(info $(shell ./bootstrap.sh $(REGISTRY_HOST) "/etc/docker/certs.d/$(REGISTRY_HOST)/ca.crt"))

.PHONY: all protoc-grpc protoc-moby go-install go-build docker-build docker-push docker-cgo docker-run
