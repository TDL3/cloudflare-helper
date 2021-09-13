FROM golang:1.17.1-alpine as builder
WORKDIR /src
COPY go.mod go.sum .
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -o cloudflare-helper .

FROM alpine:latest
WORKDIR /etc/cfh
COPY --from=builder /src/cloudflare-helper .
ENTRYPOINT [ "./cloudflare-helper" ]