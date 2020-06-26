
FROM golang:1.14.4-alpine3.12 as build-env

RUN mkdir /hello

WORKDIR /hello

# COPY the source code as the last step
COPY . .


# Build the binary
RUN go build -o /go/bin/hello


# Start the program
ENTRYPOINT ["/go/bin/hello"]
