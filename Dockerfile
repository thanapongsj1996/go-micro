FROM golang:1.21.0-alpine AS build-base

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go test --tags=unit -v ./...

RUN go build -o ./out/go-app cmd/main.go

FROM alpine:3.16.2

WORKDIR /app

COPY --from=build-base /app/out/go-app /app/go-app

CMD ["/app/go-app"]
