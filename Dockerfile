### Builder

FROM golang:1.19-alpine3.17 as builder

RUN apk --no-cache update
RUN apk add --no-cache git clang llvm make gcc protobuf

RUN mkdir /app

WORKDIR /app

COPY . .

WORKDIR /app/src

RUN go install github.com/golang/protobuf/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go build -o custom-collector

### Make executable image

FROM alpine:3.18 as custom-collector

RUN echo "@community http://dl-cdn.alpinelinux.org/alpine/edge/community" | tee -a /etc/apk/repositories

RUN apk --no-cache update
RUN apk add apparmor@community apparmor-utils@community bash

COPY --from=builder /app/src/custom-collector .

CMD ["./custom-collector"]