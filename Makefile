
# Go parameters
DOCKER_USERNAME=mchmarny
GCP_PROJECT_NAME=mchmarny-dev
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
	bin/$(BINARY_NAME) --port 8888

deps:
	go get -u github.com/tools/godep
	godep restore

gcr:
	gcloud container builds submit --tag gcr.io/$(GCP_PROJECT_NAME)/$(BINARY_NAME):latest" .

docker:
	docker build -t server-starter .
	docker tag $(BINARY_NAME):latest $(DOCKER_USERNAME)/$(BINARY_NAME):latest
	docker push $(DOCKER_USERNAME)/$(BINARY_NAME):latest"


