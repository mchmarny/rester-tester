FROM golang:1.10.1 as builder

# install ffmpeg
RUN apt-get update && apt-get install -y ffmpeg

# if app includes templates and static resources
# it may be easier to run out of the source folder 
ENV SRC_DIR=/go/src/github.com/mchmarny/rester-tester/

WORKDIR $SRC_DIR
ADD . $SRC_DIR
RUN cd $SRC_DIR

# restore to pinnned versions of dependancies 
RUN go get github.com/golang/dep/cmd/dep
RUN dep ensure

# build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o rester-tester \
    -tags netgo -installsuffix netgo .

# build the clean image
FROM scratch as runner
# copy the app
COPY --from=builder /go/src/github.com/mchmarny/rester-tester/rester-tester .

ENTRYPOINT ["/rester-tester"]