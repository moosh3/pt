FROM golang:alpine as builder

ARG build
ARG version

ENV BUILD=$build
ENV VERSION=$version

WORKDIR /build 

COPY . /build

RUN go build -o pt -ldflags "-X main.Version=$VERSION -X main.Build=$BUILD"

FROM alpine

COPY --from=builder /build/pt /bin/

ENTRYPOINT ["/bin/pt"]
CMD ["help"]
