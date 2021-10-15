FROM golang:1.16 AS builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GO_BIN=/go/bin/app

WORKDIR /var/app

COPY . .

RUN make build

FROM gcr.io/distroless/base

COPY --from=builder /go/bin/app /app

ENTRYPOINT ["./app"]