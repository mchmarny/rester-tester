
# Go parameters
DOCKER_USERNAME=mchmarny
GCP_PROJECT_NAME=serverless-lab
BINARY_NAME=rester-tester

all: test
build:
	go build -o ./bin/$(BINARY_NAME) -v

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/$(BINARY_NAME) \
	-v -a --ldflags '-extldflags "-static"' -tags netgo -installsuffix netgo

test:
	go test -v ./...

clean:
	go clean
	rm -f ./bin/$(BINARY_NAME)

run: build
	bin/$(BINARY_NAME)

deps:
	go get -u github.com/tools/godep
	godep restore

gcp:
	gcloud container builds submit --tag gcr.io/$(GCP_PROJECT_NAME)/$(BINARY_NAME):latest .

run-docker:
	docker build -t $(BINARY_NAME) .
	docker run -p 8080:8080 server-starter:latest

dockerhub: docker
	docker tag $(BINARY_NAME):latest $(DOCKER_USERNAME)/$(BINARY_NAME):latest
	docker push $(DOCKER_USERNAME)/$(BINARY_NAME):latest

deploy:
	kubectl apply -f manifest.yaml