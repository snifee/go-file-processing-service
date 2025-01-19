FROM golang:1.22.11-alpine 

WORKDIR /go-file-upload-service

COPY . /go-file-upload-service

RUN ls /go-file-upload-service

RUN go mod download

RUN go build -C cmd/file-upload-service -o file-upload-service

CMD [ "cmd/file-upload-service/file-upload-service" ]