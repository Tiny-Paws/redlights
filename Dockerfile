FROM golang:alpine

WORKDIR /go/src/app
RUN apk --no-cache add build-base git bzr mercurial gcc
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /root
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN mkdir /root/website
COPY --from=0 /go/bin/redlights .
COPY --from=0 /go/src/app/website ./website
CMD ["./redlights"]
