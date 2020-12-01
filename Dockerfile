FROM golang:alpine
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build
RUN chmod +x ./gnemes
CMD ["./gnemes"]