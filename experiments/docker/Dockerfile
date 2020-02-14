FROM golang:1.13

RUN git clone https://github.com/edenhill/librdkafka.git && \
    cd librdkafka && \
     ./configure --prefix /usr && \
     make && \
     make install

# install dependencies
RUN go get github.com/confluentinc/confluent-kafka-go/kafka && go get github.com/gorilla/websocket


# copy app
ADD . /app
WORKDIR /app/experiments

# build
RUN go build -o wsClient *.go

ENTRYPOINT ["/app/experiments/wsClient"]
