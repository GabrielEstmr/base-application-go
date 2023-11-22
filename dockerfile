FROM golang:1.20.0-alpine

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app/mpindicatorgo

# We want to populate the module cache based on the go.{mod,sum} files.
COPY src/go.mod .
COPY src/go.sum .

RUN go mod download

COPY src/ .


# Build the Go app
RUN go build -o ./out/mpindicatorgo .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./out/mpindicatorgo"]