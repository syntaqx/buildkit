FROM golang:1.16.2-alpine AS builder
WORKDIR /opt

RUN apk add --no-cache ca-certificates curl bash make git
ENV GOOS=linux GOARCH=amd64

COPY go.* ./
RUN go mod download

COPY . .
RUN make generate
RUN CGO_ENABLED=0 go install -ldflags '-s -w -extldflags "-static"' ./cmd/...

FROM alpine:3

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/bin/* /usr/local/bin/

RUN addgroup -g 1000 buildkit && \
    adduser -u 1000 -G buildkit -s /bin/sh -D buildkit
USER buildkit

ENV PORT 8080
EXPOSE $PORT

CMD ["buildkit", "server"]
