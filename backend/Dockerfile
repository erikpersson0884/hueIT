FROM golang:1.13

RUN mkdir /app
WORKDIR /app

COPY ./go.mod /app/
COPY ./go.sum /app/
RUN go mod download

COPY . /app
RUN go build ./cmd/huego

FROM alpine

ENV hue_base_url