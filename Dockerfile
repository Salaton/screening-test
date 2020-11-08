## We specify the base image we need for our
## go application
FROM golang:1.14 as builder
## We create an /app directory within our
## image that will hold our application source
## files
ENV GO111MODULE=on
RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
# ADD . /app

# ARG DB_USER
# ENV DB_USER=${DB_USER}

# ARG DB_PASSWORD
# ENV DB_PASSWORD=${DB_PASSWORD}

# ARG DB_NAME
# ENV DB_NAME=${DB_NAME}

# ARG DB_PORT
# ENV DB_PORT=${DB_PORT}

# ARG DB_HOST
# ENV DB_HOST=${DB_HOST}

# ARG SSLMODE
# ENV SSLMODE=${SSLMODE}


# ARG AFRICASTALKINGUSERNAME
# ENV AFRICASTALKINGUSERNAME=${AFRICASTALKINGUSERNAME}

# ARG AFRICASTALKINGAPIKEY
# ENV AFRICASTALKINGAPIKEY=${AFRICASTALKINGUSERNAME}
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /app
# We want to populate the module cache based on the go.{mod,sum} files.
# COPY go.mod .
# COPY go.sum .
COPY go.* $D/
CMD go mod download 

## Add this go mod download command to pull in any dependencies
RUN go get -d -v ./...

COPY . /app

## we run go build to compile the binary
## executable of our Go program
RUN cd /app/ && CGO_ENABLED=0 GOOS=linux go build -v -o server github.com/Salaton/screening-test

FROM alpine:3
RUN apk add --no-cache ca-certificates 
COPY --from=builder /app/server /server
# Expose a port to the outside world to run our app
EXPOSE 8080
## Our start command which kicks off
## our newly created binary executable
CMD ["/server"]