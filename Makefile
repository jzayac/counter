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

run-flame-graph:
	@echo ":::run flame graph"
	docker run uber/go-torch --link $(CONTAINER_NAME):$(CONTAINER_NAME) -u http://$(CONTAINER_NAME):$(PORT) -p > torch.svg

test-request:
	@echo ":::test single request"
	curl localhost:8080 

benchmark:
	@echo ":::bechmark and test data race"
	docker run -it --rm --link $(CONTAINER_NAME) $(IMAGE_NAME):$(BASE_TAG) go-wrk -d 5 http://$(CONTAINER_NAME):8080

