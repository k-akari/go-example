FROM golang:1.18-alpine

ENV APP_HOME /app
RUN mkdir $APP_HOME
WORKDIR $APP_HOME
ADD ../app $APP_HOME

#RUN go get github.com/lib/pq