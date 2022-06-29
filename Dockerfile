FROM golang:1.18-alpine as builder

WORKDIR /src/
RUN apk add --no-cache alpine-sdk
# TODO: It actually copies everything in project folder to builder. Fix later.
COPY . .
RUN ls /src/; go mod download; CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o /bin/nolocks /src/cmd/main.go

FROM alpine:3.14

COPY --from=builder /bin/nolocks /nolocks
COPY config.yml /
RUN apk add --no-cache ca-certificates
CMD ["./nolocks"]