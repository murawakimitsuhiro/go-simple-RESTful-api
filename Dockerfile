FROM golang:latest

ENV TZ Asia/Tokyo
RUN apt-get update \
  && apt-get install -y tzdata \
  && rm -rf /var/lib/apt/lists/* \
  && echo "${TZ}" > /etc/timezone \
  && rm /etc/localtime \
  && ln -s /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
  && dpkg-reconfigure -f noninteractive tzdata

ENV GOPATH /go

RUN apt-get update && \
    apt-get upgrade -y

RUN mkdir /go/src/go-simple-RESTful-api
WORKDIR /go/src/go-simple-RESTful-api
COPY ./ /go/src/go-simple-RESTful-api

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -v -update

CMD ["go", "run", "main.go"]

EXPOSE 8005
