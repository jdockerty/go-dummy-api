# STAGE ONE
FROM golang:onbuild as builder

LABEL description="A dummy API that wraps a 'jsonplaceholder' for testing/learning various services."

WORKDIR /build
# Copy appropriate golang files into the working directory
# This ignores the files contained within the .dockerignore file
COPY . . 

# Building the .go file ready for Alpine Linux container
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dummy-api .

# STAGE TWO

FROM alpine:latest

WORKDIR /api

# Using previous stage to grab binary file into current stage
COPY --from=builder /build/dummy-api .

# For documentation only, this lets the reader know that the application exposes port 8080
# which can then be mapped from the container port to the host with the command
# docker container run -p 8080:<host_port> <image>
EXPOSE 8080

ENTRYPOINT [ "./dummy-api" ]
