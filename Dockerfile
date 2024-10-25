# Step 1: Modules caching
FROM golang:1.23.2-alpine as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.23.2-alpine as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates
RUN CGO_ENABLED=0 GOOS=linux  \
    go build  -o /bin/app ./cmd/main.go

# Step 3: Final
FROM scratch
COPY --from=builder bin/app /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/app"]
