# syntax=docker/dockerfile:1

FROM golang:1.23.1

# Set destination for COPY
RUN apt-get update && apt-get install -y uuid-runtime
WORKDIR /app
COPY . .
# Download Go modules
# COPY go.mod  ./
RUN cd src
RUN pwd
RUN go mod init opendrm

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
# COPY *.go ./

# Build
RUN cd src && ls -al

RUN cd src && ls -al && CGO_ENABLED=0 GOOS=linux go build -o /opendrm

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8090

# Run
CMD ["/opendrm"]
