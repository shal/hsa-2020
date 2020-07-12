FROM golang:1.14-alpine AS build

ENV GO111MODULE=on

WORKDIR /go/src/app

RUN apk add bash ca-certificates git gcc g++ libc-dev curl make

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN BLDDIR=/go/bin make build

FROM alpine

RUN apk update && apk upgrade && apk add curl

WORKDIR /app

COPY --from=build /go/bin/ ./
COPY ./config ./config

CMD ["./server"]
