# Build the manager binary
FROM golang:1.14 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/
COPY pkg/ pkg/

# Build
RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -a -o manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM ubuntu:16.04
RUN apt-get update -y && apt-get install ca-certificates -y && apt-get install curl -y
WORKDIR /
COPY --from=builder /workspace/manager .

RUN useradd -ms /bin/bash devtron
RUN chown -R devtron:devtron ./manager
USER devtron

ENTRYPOINT ["/manager"]
