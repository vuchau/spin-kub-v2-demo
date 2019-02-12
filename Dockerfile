# Accept the Go version for the image to be set as a build argument.
# Default to Go 1.11
ARG GO_VERSION=1.11.4

# First stage: build the executable.
FROM golang:${GO_VERSION}-alpine AS builder

# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
RUN apk add --no-cache ca-certificates git

# Set the environment variables for the go command:
ENV CGO_ENABLED=0 GOFLAGS=-mod=vendor GOOS=linux GOARCH=amd64

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Import the code from the context.
COPY ./ ./

# Build the executable to `/webapp`. Mark the build as statically linked.
RUN go build -installsuffix 'static' -o /hello .

# Final stage: the running container.
FROM scratch

# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Import the compiled executable from the second stage.
COPY --from=builder /hello /hello

EXPOSE 8000

# Executable
ENTRYPOINT ["/hello"]
