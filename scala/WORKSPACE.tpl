rules_scala_version = "eeee4679d07eed3a12666361aecbc556047a4f17"

http_archive(
    name = "io_bazel_rules_scala",
    strip_prefix = "rules_scala-%s" % rules_scala_version,
    type = "zip",
    url = "https://github.com/bazelbuild/rules_scala/archive/%s.zip" % rules_scala_version,
)

load("@io_bazel_rules_scala//scala:scala.bzl", "scala_repositories")

scala_repositories()

load("@io_bazel_rules_scala//scala:toolchains.bzl", "scala_register_toolchains")

scala_register_toolchains()

# Uncomment to define your own Scala toolchain
# See: https://github.com/bazelbuild/rules_scala#b-defining-your-own-scala_toolchain-requires-2-steps
# register_toolchains("//toolchains:my_scala_toolchain")
