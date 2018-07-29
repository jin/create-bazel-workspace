load("@io_bazel_rules_scala//scala:scala.bzl", "scala_library", "scala_binary", "scala_test")

# Uncomment to define your own Scala toolchain
# See: https://github.com/bazelbuild/rules_scala#b-defining-your-own-scala_toolchain-requires-2-steps

# load("@io_bazel_rules_scala//scala:scala_toolchain.bzl", "scala_toolchain")

# scala_toolchain(
#     name = "my_toolchain_impl",
#     scalacopts = ["-Ywarn-unused"],
#     visibility = ["//visibility:public"]
# )

# toolchain(
#     name = "my_scala_toolchain",
#     toolchain_type = "@io_bazel_rules_scala//scala:toolchain_type",
#     toolchain = "my_toolchain_impl",
#     visibility = ["//visibility:public"]
# )
