FROM golang:latest
WORKDIR /go/src/github.com/joram/jsnek/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o jsnek .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/joram/jsnek/jsnek .
CMD ["./jsnek"]