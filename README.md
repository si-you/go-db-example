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
