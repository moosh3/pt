# Build Step
FROM golang:1.10 AS build

# Prerequisites and vendoring
RUN mkdir -p $GOPATH/src/path/to/your/project/
ADD . $GOPATH/src/path/to/your/project/
WORKDIR $GOPATH/src/path/to/your/project/
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -vendor-only

# Build
ARG build
ARG version
RUN CGO_ENABLED=0 go build -ldflags="-s -w -X main.Version=${version} -X main.Build=${build}" -o myprogram
RUN cp myprogram /

# Final Step
FROM alpine

# Base packages
RUN apk update
RUN apk upgrade
RUN apk add ca-certificates && update-ca-certificates
RUN apk add --update tzdata
RUN rm -rf /var/cache/apk/*

# Copy binary from build step
COPY --from=build /myprogram /home/
# Define timezone
ENV TZ=Europe/Paris

# Define the ENTRYPOINT
WORKDIR /home
ENTRYPOINT ./myprogram

# Document that the service listens on port 8080.
EXPOSE 8080
