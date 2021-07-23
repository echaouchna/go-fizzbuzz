FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /fizzbuzz-server

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /fizzbuzz-server /fizzbuzz-server

EXPOSE 10000

USER nonroot:nonroot

ENTRYPOINT ["/fizzbuzz-server"]