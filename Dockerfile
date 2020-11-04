## We specify the base image we need for our
## go application
FROM golang:alpine as builder
## We create an /app directory within our
## image that will hold our application source
## files
ENV GO111MODULE=on
RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD . /app
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /app
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

## Add this go mod download command to pull in any dependencies
RUN go get -d -v ./...

COPY . .

## we run go build to compile the binary
## executable of our Go program
RUN go build -o main .
# Expose a port to the outside world to run our app
EXPOSE 8080
## Our start command which kicks off
## our newly created binary executable
CMD ["/app/main"]