BINARY_NAME = top
IMAGE = smvfal/top
TAG = latest

install:
	kubectl apply -f kubernetes/rbac.yml
	kubectl apply -f kubernetes/deployment.yml

publish: docker-build docker-push

docker-build:
	DOCKER_BUILDKIT=1 docker build -t ${IMAGE}:${TAG} .

docker-push:
	docker push ${IMAGE}:${TAG}

vendor:
	go mod vendor -v

build:
	go build -o bin/${BINARY_NAME}

clean:
	rm -rf vendor/ bin/