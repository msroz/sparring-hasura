FROM golang:1.16-alpine3.13 as base

WORKDIR /app/go/base

COPY go.mod .
COPY go.sum . 

RUN go mod download
RUN go get -u github.com/cosmtrek/air

COPY . .

FROM golang:1.16-alpine3.13 as builder

WORKDIR /app/go/builder

COPY --from=base /app/go/src /app/go/builder
RUN CGO_ENABLED=0 go build main.go

FROM scratch as runner

WORKDIR /app/go/src

RUN apk add --no-cache
COPY --from=builder /app/go/builder/main /app/go/src/main

EXPOSE 9876

CMD ["/app/go/src/main"]
