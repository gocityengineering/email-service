FROM golang:1.17 as builder
WORKDIR /go/src/github.com/gocityengineering/email-service
ADD . ./
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GO111MODULE on
RUN \
  go get && \
  go vet && \
  go test -v ./... && \
  go build

FROM ubuntu:21.04
WORKDIR /app/
RUN groupadd app && useradd -g app app
COPY --from=builder /go/src/github.com/gocityengineering/email-service/email-service /usr/local/bin/email-service
USER app
CMD ["email-service"]
