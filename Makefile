# BASE_TAG=$(shell git rev-parse --short HEAD)
BASE_TAG=latest
IMAGE_NAME=rabbyt/counter
PORT=8080
CONTAINER_NAME=counter

build-test:
	@echo ":::build develop"
	docker build --rm -t $(IMAGE_NAME):$(BASE_TAG) -f Dockerfile-develop .

run-test-image:
	@echo ":::run test image"
	docker run -d --rm --name $(CONTAINER_NAME) -p $(PORT):$(PORT) $(IMAGE_NAME):$(BASE_TAG)

test-request:
	@echo ":::test single request"
	curl localhost:8080

benchmark:
	@echo ":::bechmark and test data race"
	docker run -it --rm --link $(CONTAINER_NAME) $(IMAGE_NAME):$(BASE_TAG) go-wrk -d 5 http://$(CONTAINER_NAME):8080

