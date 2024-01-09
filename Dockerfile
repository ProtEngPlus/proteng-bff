#build stage
FROM golang:1.21 AS builder
ARG ARCH=amd64
WORKDIR /go/src/proteng-bff
ADD . .
RUN go get proteng-bff
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} go build -a -ldflags '-extldflags "-static"' -o app ./main.go

#final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/proteng-bff/app .
CMD ["./app"]

EXPOSE 8080