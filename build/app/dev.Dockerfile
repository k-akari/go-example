FROM golang:1.18-alpine

ENV APP_HOME /go/src/project
RUN mkdir $APP_HOME
WORKDIR $APP_HOME
ADD ../.. $APP_HOME

RUN apk update && apk add git openssh gcc libc-dev

RUN go install -v golang.org/x/tools/gopls@latest \
  honnef.co/go/tools/cmd/staticcheck@latest
