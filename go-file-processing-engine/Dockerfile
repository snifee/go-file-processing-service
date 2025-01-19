FROM golang:1.22.11-alpine 

WORKDIR /go-file-processing-engine

COPY . /go-file-processing-engine

RUN ls /go-file-processing-engine

RUN go mod download

RUN go build -C cmd/file-processing-engine -o file-processing-engine

CMD [ "cmd/file-processing-engine/file-processing-engine" ]