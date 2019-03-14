FROM golang:1.11.1-alpine
RUN apk add --update git; \
    mkdir -p ${GOPATH}/go-rest-api; \
    go get -u github.com/D3n0Duz/db-service 
WORKDIR ${GOPATH}/go-rest-api/

#RUN git clone https://github.com/D3n0Duz/db-service.git
#RUN go build -o go-rest-api ./db-service

COPY . ${GOPATH}/go-rest-api/
RUN go build -o go-rest-api .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=0 /go/go-rest-api/go-rest-api .
#COPY --from=0 /go/go-rest-api/db-service/config.toml .
COPY --from=0 /go/go-rest-api/config.toml .
RUN ls go-rest-api
EXPOSE 3000

CMD [ "/app/go-rest-api" ]