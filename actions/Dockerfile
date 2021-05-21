FROM golang:1.16-alpine3.13 as builder
WORKDIR /go/src
COPY . .

RUN apk update && \
      apk add --no-cache git && \
      go build -o app && \
      GO111MODULE=off go get github.com/oxequa/realize

FROM scratch as runner

WORKDIR /app

COPY --from=builder /go/src/app .

RUN addgroup go \
    && adduser -D -G go go \
    && chown -R go:go /app/app

CMD ["./app"]
