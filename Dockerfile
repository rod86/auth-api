FROM golang:1.12-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh curl make

WORKDIR /go/src/app

# Go lang packages
RUN go get github.com/cespare/reflex

# Setup Go Modules
ENV GO111MODULE=on
COPY ./code ./
RUN make install

EXPOSE 8000

CMD ["make", "watch"]

