# we need this image
FROM golang:1.5.3

RUN curl -sSL https://get.docker.com/ | sh
VOLUME /var/lib/docker

RUN go get "github.com/aws/aws-sdk-go" \
           "github.com/jasonlvhit/gocron" \
           "github.com/go-ini/ini" \
           "github.com/jmespath/go-jmespath"

ENV APP_HOME=${GOPATH}/src/github.com/lunchiatto/backuper
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME
ADD . $APP_HOME

# Build dependencies and the ws binary
RUN go build
