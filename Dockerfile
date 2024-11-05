FROM golang:latest AS builder
LABEL authors="drawiin"
WORKDIR /app
COPY . .
# Set GOOS to linux to build a binary for linux
# Disable CGO for static linking
# Add ldflags to reduce binary size
# -o flag to specify the output binary name
RUN GOOS=linux CGO_ENABLED=0 go build -o runner cmd/..

# Copy the binary to the scratch image, this way we have a minimal
# image because we don't need the whole OS neither the golang image
FROM scratch
COPY --from=builder /app/runner .
# Use ENTRYPOINT to allow passing arguments to the binary
ENTRYPOINT ["./runner"]