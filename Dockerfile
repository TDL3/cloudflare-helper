FROM golang:1.16.2-alpine as build
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -o cloudflare-helper .

FROM alpine:latest
WORKDIR /etc/cfh
COPY --from=build /src/cloudflare-helper .
ENTRYPOINT [ "./cloudflare-helper" ]