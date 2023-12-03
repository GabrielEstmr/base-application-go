FROM golang:1.21.4-alpine

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app/base-application-go

# We want to populate the module cache based on the go.{mod,sum} files.
COPY src/go.mod .
COPY src/go.sum .

RUN go mod download

COPY src/ .


# Build the Go app
RUN go build -o ./out/base-application-go .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./out/base-application-go"]