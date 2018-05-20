# go-db-example.

A minimal example shows how to use various dbs in Go + Bazel.

*Redis example*
```
bazel run examples/redis:redis_example
```

*Mongo example*
```
bazel run examples/mongo:mongo_example
```

*Kafka example*
```
bazel run examples/kafka:kafka_example
```

To build a go_image
```
bazel run -c opt examples/kafka:kafka_example_image
```

*Confluent kafka example*

You need to install librdkafka first.
Tip) ./configure --prefix=/usr


```
bazel run --linkopt=-lrdkafka examples/confluent_kafka:confluent_kafka_example
```

To build a container_image, run third_party/librdkafka/build.sh to prepare a base docker image as a tar file. Then,

```
sudo bazel run --linkopt=-lrdkafka examples/confluent_kafka:confluent_kafka_example_image
```

TODO(si-you): Make librdkafka image compact.. It's too large now.
