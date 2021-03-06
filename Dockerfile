FROM golang:1.13 as builder

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o restdb

FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY --from=builder /workspace/restdb .
COPY config.yml /config.yml

USER nonroot:nonroot

ENTRYPOINT ["./restdb"]
