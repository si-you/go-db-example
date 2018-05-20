FROM ubuntu:16.04

RUN apt-get update && \
    apt-get -y upgrade 
RUN apt-get install -y \
      make wget libpthread-stubs0-dev build-essential
RUN rm -rf /var/lib/apt/lists/*

ADD librdkafka /tmp/workspace/librdkafka
WORKDIR /tmp/workspace/librdkafka
RUN ./configure --prefix=/usr
RUN make
RUN make install

RUN wget https://storage.googleapis.com/golang/go1.10.2.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.10.2.linux-amd64.tar.gz && \
    rm -rf /go1.10.2.linux-amd64.tar.gz
ENV PATH $PATH:/usr/local/go/bin

WORKDIR /workspace
