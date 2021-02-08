FROM golang:alpine as builder

# Application Directory
RUN mkdir /app
WORKDIR /app

# First handle dependencies as those probably are more stable than rest of codebase
COPY ./go.mod /app/
COPY ./go.sum /app/
RUN go mod download

ENV HUE_BASE_URL=""

# Copy source and build app
COPY . /app
RUN go build -o huego ./cmd/huego/main.go

FROM alpine

# Copy over the app from the builder image
COPY --from=builder /app/huego /huego
COPY --from=builder /app/config.json /

EXPOSE 8080

ENTRYPOINT ["/huego"]