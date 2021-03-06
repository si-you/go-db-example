git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "167cb55932c4a415edf839f750bdd7366f2f1613",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains", "go_repository")

# Repos for radix.
go_repository(
  name = "com_github_mediocregopher_radix_v2",
  importpath = "github.com/mediocregopher/radix.v2",
  commit = "596a3ed684d9390f4831d6800dffd6a6067933d5",
)

# Repos for mongo.
go_repository(
  name = "com_github_mongodb_mongo_go_driver",
  importpath = "github.com/mongodb/mongo-go-driver",
  commit = "ff5ad991913ea56c4609391ff190eaf83edf339b",
)

go_repository(
  name = "com_github_go_stack_stack",
  importpath = "github.com/go-stack/stack",
  commit = "259ab82a6cad3992b4e21ff5cac294ccb06474bc",
)

go_repository(
  name = "org_golang_x_sync",
  importpath = "golang.org/x/sync",
  commit = "1d60e4601c6fd243af51cc01ddf169918a5407ca",
)

go_repository(
  name = "org_golang_x_crypto",
  importpath = "golang.org/x/crypto",
  commit = "2d027ae1dddd4694d54f7a8b6cbe78dca8720226",
)

# Repos for kafka-go
go_repository(
  name = "com_github_segmentio_kafka_go",
  importpath = "github.com/segmentio/kafka-go",
  commit = "44a19ca9cf5925ba3c5153dd0d9e5b44e8dfd717",
)

# Repos for confluent-kafka
# WARNING: This repo requires librdkafka installed in your system.
go_repository(
  name = "com_github_confluentinc_confluent_kafka_go",
  importpath = "github.com/confluentinc/confluent-kafka-go",
  commit = "5e4d04e05fc319ce5996a867aafef29059f26862",
)

go_rules_dependencies()
go_register_toolchains()

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "6dede2c65ce86289969b907f343a1382d33c14fbce5e30dd17bb59bb55bb6593",
    strip_prefix = "rules_docker-0.4.0",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/v0.4.0.tar.gz"],
)


load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull", "container_load",
    container_repositories = "repositories",
)

# This is NOT needed when going through the language lang_image
# "repositories" function(s).
container_repositories()

container_load(
  name = "librdkafka",
  file = "//third_party/librdkafka:librdkafka.tar",
)

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()
