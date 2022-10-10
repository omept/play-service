####################################################################
# Builder Stage                                                    #
####################################################################
# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:alpine AS builder

LABEL MAINTAINER="georgetheprogrammer@gmail.com"

# Create WORKDIR using project's root directory
WORKDIR /go/src/github.com/ong-gtp/play-service
# Copy the local package files to the container's workspace
# in the above created WORKDIR

COPY . .
RUN go get github.com/gorilla/mux
RUN go get github.com/go-kit/kit
RUN go get github.com/go-kit/log
RUN go mod tidy
# Build the API service inside the container
RUN go build -o playsv


#####################################################################
# Final Stage                                                       #
#####################################################################
# Pull golang alpine image (very small image, with minimum needed to run Go)
FROM alpine:3.16

RUN apk update \
    && apk upgrade

# Create WORKDIR
WORKDIR /app

# Copy app binary from the Builder stage image
COPY --from=builder /go/src/github.com/ong-gtp/play-service/playsv . 

# Run the playsv command by default when the container starts.
CMD ["./playsv"]

# Document that the service uses port 8080
EXPOSE 8082

