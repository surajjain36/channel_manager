# Stage 1
FROM golang:1.12.8 AS builder
# Creating working directory
RUN mkdir -p  /go/src/github.com/surajjain36
# Copying source code to repository
COPY   .  /go/src/github.com/surajjain36/channel_manager
WORKDIR /go/src/github.com/surajjain36/channel_manager
# Installing ca certificates
RUN apt-get update && apt-get install --no-install-recommends -y ca-certificates && rm -rf /var/lib/apt/lists/*
ENV GO111MODULE=on

# Creating go binary
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o channel_manager
# Stage 2
FROM alpine
RUN apk add --no-cache openssh
# Copy ca certificates from builder
COPY --from=builder  /etc/ssl/certs /etc/ssl/certs
# Copy our static executable and dependencies from builder
COPY --from=builder /go/src/github.com/surajjain36/channel_manager/channel_manager /
COPY --from=builder /go/src/github.com/surajjain36/channel_manager/config.yml  /

# Exposing port
EXPOSE 3000
# Run the hannel_manager binary.
ENTRYPOINT ["/channel_manager"]