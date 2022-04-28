FROM golang:1.18-alpine

ENV APP_HOME /go/src/project
RUN mkdir $APP_HOME
WORKDIR $APP_HOME
ADD .. $APP_HOME

RUN apk update && apk add git

# install go tools（自動補完等に必要なツールをコンテナにインストール）
RUN go get github.com/uudashr/gopkgs/v2/cmd/gopkgs \
  github.com/ramya-rao-a/go-outline \
  github.com/nsf/gocode \
  github.com/acroca/go-symbols \
  github.com/fatih/gomodifytags \
  github.com/josharian/impl \
  github.com/haya14busa/goplay/cmd/goplay \
  github.com/go-delve/delve/cmd/dlv \
  golang.org/x/lint/golint \
  golang.org/x/tools/gopls \
  github.com/lib/pq