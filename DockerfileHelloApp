FROM golang:latest

WORKDIR $GOPATH/src/hello-app
COPY . .
RUN go get -d -v
RUN go install 

FROM alpine:latest
COPY --from=0 /go/bin/hello-app .
ENV PORT 3000

CMD ["./hello-app"]