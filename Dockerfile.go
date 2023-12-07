# Start from golang base image
FROM golang:1.21-alpine as builder

ENV GO111MODULE=on

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the working directory inside the container
WORKDIR /go/src/advent_of_code

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o /go/bin/advent_app ./cmd

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates


# Copy only the necessary files from the build stage
COPY --from=builder /go/bin/advent_app /advent_app
COPY --from=builder /go/src/advent_of_code/inputfiles /inputfiles

# Expose port 8084 to the outside world
EXPOSE 8084

# Command to run the executable
CMD ["./advent_app"]