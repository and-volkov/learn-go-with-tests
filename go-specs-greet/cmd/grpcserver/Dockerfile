FROM golang:1.24.1-alpine AS deps

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

FROM golang:1.24.1-alpine AS build

WORKDIR /app

COPY --from=deps /go/pkg /go/pkg
COPY --from=deps /app/go.mod /app/go.sum

COPY . .

RUN go build -o svr cmd/grpcserver/*.go

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/svr .

EXPOSE 50051
CMD [ "./svr" ]
