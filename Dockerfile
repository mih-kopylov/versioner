FROM golang:1.21-alpine AS builder
WORKDIR /

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build

FROM alpine/git
WORKDIR /app

COPY --from=builder /versioner /bin/versioner

RUN apk add --update bash && apk add --update git-lfs && rm -rf /var/cache/apk/*

ENTRYPOINT ["bash"]
