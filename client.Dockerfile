FROM golang:alpine as build

RUN apk add ca-certificates git

WORKDIR /app

COPY . ./
RUN  go mod download


RUN go test -v ./...

RUN cd /app/cmd/client && \
    go build -o /client


FROM alpine:latest

COPY --from=build /client /client

WORKDIR /srv
ENTRYPOINT ["/client"]