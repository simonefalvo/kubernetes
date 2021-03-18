BINARY_NAME = secretreader
IMAGE = smvfal/secretreader
TAG = latest

docker-build:
	DOCKER_BUILDKIT=1 docker build -t ${IMAGE}:${TAG} .

docker-push:
	docker push ${IMAGE}:${TAG}

vendor:
	go mod vendor -v

build:
	go build -o ${BINARY_NAME}

clean:
	rm -r vendor/
	rm ${BINARY_NAME}