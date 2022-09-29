FROM golang:1.19-bullseye

# RUN apk add gcc

COPY . /hack

WORKDIR /hack

RUN go build -o pricewatcher .

ENTRYPOINT ["/hack/pricewatcher"]