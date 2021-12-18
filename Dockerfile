FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on 
RUN go get github.com/Afaryz/goDockerWebAPI
RUN cd /build && git clone https://github.com/Afaryz/goDockerWebAPI.git
RUN cd /build/goDockerWebAPI/app && go build -o main
EXPOSE 8080

ENTRYPOINT [ "/build/goDockerWebAPI/app/main" ]