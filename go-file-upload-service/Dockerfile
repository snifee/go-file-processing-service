FROM golang:1.22.11-alpine 

WORKDIR /go-file-upload-service

COPY . /go-file-upload-service

RUN ls /go-file-upload-service

RUN go mod download

RUN go build -o file-upload-service

CMD [ "file-upload-service" ]