FROM golang:1.17 as builder

WORKDIR /workspace
COPY go.mod go.mod
RUN go mod download

COPY main.go main.go
COPY images/ images/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o image-source main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/image-source .
USER 65532:65532

ENTRYPOINT ["/image-source"]
