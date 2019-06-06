FROM golang:1.12-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh curl

WORKDIR /go/src/app

# Go lang packages
RUN go get github.com/cespare/reflex

# Setup Go Modules
ENV GO111MODULE=on
COPY ./code/go.mod .
COPY ./code/go.sum .
RUN go mod download
COPY ./code ./

EXPOSE 8000

CMD ["reflex", "-c", "reflex.conf"]

