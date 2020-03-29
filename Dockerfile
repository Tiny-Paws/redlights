FROM golang

WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/bin/app .
CMD ["./app"]
