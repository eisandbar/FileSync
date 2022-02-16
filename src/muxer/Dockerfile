FROM golang:1.17
WORKDIR /filesync
COPY go.mod ./
RUN go mod download
USER root
RUN apt-get update
RUN apt-get install ffmpeg -y
COPY . .
RUN go build -o ./bin/
CMD ["./bin/muxer"]