# syntax=docker/dockerfile:1

FROM golang:latest AS build
WORKDIR /src/gin
COPY go.mod ./
RUN go get -v github.com/gin-gonic/gin
COPY main.go ./
RUN CGO_ENABLED=0 go build -o gin-server

FROM alpine:latest  
WORKDIR /root/
COPY --from=build /src/gin/gin-server ./
CMD ["./gin-server"]