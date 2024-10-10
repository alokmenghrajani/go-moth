FROM golang:1.21.4

EXPOSE 8000

COPY . /go-moth
RUN cd /go-moth && \
    go build .

ENTRYPOINT ["/go-moth/go-moth"]

