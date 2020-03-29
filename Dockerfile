FROM golang

WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /root
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY --from=0 /go/bin/app .
CMD ["./app"]
