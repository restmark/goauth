FROM        golang:1.8
   
# Setting up working directory
WORKDIR     /go/src/github.com/restmark/goauth
ADD         . /go/src/github.com/restmark/goauth

RUN apt-get update && apt-get install -y \
    libsasl2-dev

# Install the dependencies
RUN go get -t -v ./...

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/restmark/goauth

ENV GIN_MODE release
ENV GOAUTH_ENV prod

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/goauth

# Document that the service listens on port 4000.
EXPOSE 4000