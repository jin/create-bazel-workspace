load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "c1f52b8789218bb1542ed362c4f7de7052abcf254d865d96fb7ba6d44bc15ee3",
    urls = [
        "https://github.com/bazelbuild/rules_go/releases/download/0.12.0/rules_go-0.12.0.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "ddedc7aaeb61f2654d7d7d4fd7940052ea992ccdb031b8f9797ed143ac7e8d43",
    urls = [
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.12.0/bazel-gazelle-0.12.0.tar.gz",
    ],
)

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_rules_dependencies",
    "go_register_toolchains",
)

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "com_github_geertjohan_go_rice",
    commit = "c02ca9a983da5807ddf7d796784928f5be4afd09",
    importpath = "github.com/GeertJohan/go.rice",
)

go_repository(
    name = "com_github_kardianos_osext",
    commit = "ae77be60afb1dcacde03767a8c37337fad28ac14",
    importpath = "github.com/kardianos/osext",
)

go_repository(
    name = "com_github_daaku_go_zipexe",
    commit = "a5fe2436ffcb3236e175e5149162b41cd28bd27d",
    importpath = "github.com/daaku/go.zipexe",
)
