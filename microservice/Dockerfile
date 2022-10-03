FROM golang:1.19 AS base

RUN apt-get update \
  && apt-get -y install zip make openssh-client ca-certificates \
  && update-ca-certificates

FROM base AS builder
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod tidy

COPY Makefile .
COPY . ./
RUN make tidy
RUN make build

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/microservice /microservice

EXPOSE 8090

CMD ["/microservice", "serve"]