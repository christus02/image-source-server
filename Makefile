IMG ?= raghulc/image-server:latest

all: build

build:
	go build -o bin/image-source main.go

docker-build:
	docker build -t ${IMG} .

docker-push:
	docker push ${IMG}

docker: docker-build docker-push
