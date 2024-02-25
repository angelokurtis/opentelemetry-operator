# Build Stage using golang image as the base
FROM golang:1.22.0 as builder

# Define build arguments
ARG TARGETOS
ARG TARGETARCH

# Set environment variables
ENV CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH}

WORKDIR /workspace

# Copy only the go.mod and go.sum files to leverage Docker layer caching
COPY go.mod go.mod
COPY go.sum go.sum

# Download Go modules to cache them for future builds
RUN go mod download

# Copy the application source code to the container
COPY apis/ apis/
COPY controllers/ controllers/
COPY internal/ internal/
COPY pkg/ pkg/
COPY main.go main.go

# Build the Go application
RUN go build -ldflags "-s -w" -a -o bin/opentelemetry-operator

# Get CA certificates from alpine package repo
FROM alpine:3.19 as certificates

RUN apk --no-cache add ca-certificates

######## Start a new stage from scratch #######
FROM scratch

ARG TARGETARCH

ENV USER=kurtis

WORKDIR /

# Copy the certs from Alpine
COPY --from=certificates /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Copy binary built on the host
COPY --from=builder /workspace/bin/opentelemetry-operator manager

USER 65532:65532

ENTRYPOINT ["/manager"]
