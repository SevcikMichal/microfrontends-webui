FROM node:latest AS webui

ARG PUPPETEER_SKIP_DOWNLOAD=true
RUN mkdir /app
WORKDIR /app

COPY web-ui/package.json .
COPY web-ui/package-lock.json .
RUN npm install 

COPY web-ui .

ARG BUILD_ENV=build.prod

RUN npm run $BUILD_ENV

FROM golang:1.20 as builder
ARG TARGETOS
ARG TARGETARCH

WORKDIR /workspace

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY model/ model/
COPY router/ router/
COPY server/ server/
COPY configuration/ configuration/
COPY main.go main.go

# Build
# the GOARCH has not a default value to allow the binary be built according to the host where the command
# was called. For example, if we call make docker-build in a local env which has the Apple Silicon M1 SO
# the docker BUILDPLATFORM arg will be linux/arm64 when for Apple x86 it will be linux/amd64. Therefore,
# by leaving it empty we can ensure that the container and binary shipped on it will have the same platform.
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -a -o web-ui-server main.go

# Use distroless as minimal base image to package the web-ui-server binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/web-ui-server web-ui-server
COPY --from=webui /app/www/ /web-ui/www/
USER 65532:65532

ENTRYPOINT ["/web-ui-server"]
