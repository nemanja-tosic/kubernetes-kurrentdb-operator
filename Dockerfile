# Build the manager binary
FROM --platform=$BUILDPLATFORM golang:1.25 AS builder
ARG TARGETOS
ARG TARGETARCH

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
# Download dependencies (no go.sum required here)
RUN go mod download

# Copy the go source
COPY cmd/ cmd/
COPY api/ api/
COPY controllers/ controllers/

# Resolve and tidy modules inside the builder (ensures go.sum is generated)
RUN go mod tidy

# Build
# Build for linux using the container's native architecture to avoid shell/arg expansion issues across hosts
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -a -o manager cmd/main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/manager .
USER 65532:65532

ENTRYPOINT ["/manager"]
