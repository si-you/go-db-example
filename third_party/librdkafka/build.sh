#!/bin/sh

docker build -t librdkafka:0.11.4 .
docker save librdkafka:0.11.4 > librdkafka.tar
