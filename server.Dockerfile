FROM golang:alpine as build

RUN apk add ca-certificates git

WORKDIR /app

COPY . ./
RUN  go mod download


RUN go test -v ./...

RUN cd /app/cmd/server && \
    go build -o /server


FROM alpine:latest

COPY --from=build /server /server

WORKDIR /srv
ENTRYPOINT ["/server"]