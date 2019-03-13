FROM golang:latest 

WORKDIR $GOPATH/src/
COPY . .
RUN go get -d -v

RUN ls
RUN go build -o app.go


EXPOSE 3000

CMD ["./app.go"]

