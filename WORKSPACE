load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name="io_bazel_rules_go",
    sha256="c1f52b8789218bb1542ed362c4f7de7052abcf254d865d96fb7ba6d44bc15ee3",
    urls=[
        "https://github.com/bazelbuild/rules_go/releases/download/0.12.0/rules_go-0.12.0.tar.gz"
    ],
)

http_archive(
    name="bazel_gazelle",
    sha256="ddedc7aaeb61f2654d7d7d4fd7940052ea992ccdb031b8f9797ed143ac7e8d43",
    urls=[
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.12.0/bazel-gazelle-0.12.0.tar.gz"
    ],
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies",
     "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()
