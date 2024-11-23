ARG GO_VERSION=1.22.1

FROM golang:${GO_VERSION} AS build

WORKDIR /opt/gogotravelbot

COPY . ./

RUN go mod download \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gogotravelbot cmd/bot/main.go

FROM istio/distroless:latest AS run

WORKDIR /opt/gogotravelbot

COPY --from=build /opt/gogotravelbot/gogotravelbot .

CMD ["./gogotravelbot"]