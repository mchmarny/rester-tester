FROM golang
MAINTAINER Mark Chmarny <mark@chmarny.com>

# get ffmpeg
RUN apk add --no-cache ffmpeg

# if app includes templates and static resources
# it may be easier to run out of the source folder 
ENV SRC_DIR=/go/src/github.com/mchmarny/rester-tester/

WORKDIR $SRC_DIR
ADD . $SRC_DIR
RUN cd $SRC_DIR

# restore to pinnned versions of dependancies 
RUN go get github.com/tools/godep
RUN godep restore

RUN go build

ENTRYPOINT $SRC_DIR/rester-tester
