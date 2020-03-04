FROM golang:alpine as builder

WORKDIR /go/src/go-hello-world
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o app

FROM alpine
RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/go-hello-world/app /app

CMD ["/app"]