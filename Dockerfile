FROM golang:1.20-alpine3.18 AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/app/main ./cmd/app

FROM alpine:3.18.0

RUN apk --no-cache add tzdata

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /build/cmd/app/main /
COPY --from=builder /build/configs/ /configs/
COPY --from=builder /build/web /web

EXPOSE 8080

ENTRYPOINT ["./main"]