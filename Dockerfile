FROM golang:1.11 AS builder
#WORKDIR /data
#WORKDIR $GOPATH/src
WORKDIR $GOPATH/src/web-yamaha
COPY *.go .
#RUN go env GOPATH
RUN go get ./...
#RUN go mod download
#RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
COPY --from=builder /data/app .
COPY static .
EXPOSE 8080/tcp
ENTRYPOINT [ "./app" ]
