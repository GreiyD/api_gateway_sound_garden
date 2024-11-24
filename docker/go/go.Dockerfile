FROM golang:1.23 AS builder

ARG WORKDIR
WORKDIR ${WORKDIR}

COPY . .

ENV GOCACHE=/go/cache

RUN mkdir -p /go/cache

RUN go build -o api-gateway ./cmd/api-gateway/main.go

FROM alpine:3.15

RUN apk add --no-cache libc6-compat

ARG WORKDIR

COPY --from=builder ${WORKDIR}/api-gateway .

ENV GOCACHE=/go/cache

CMD ["./api-gateway"]

EXPOSE 5000