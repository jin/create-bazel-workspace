load("@build_bazel_rules_android//android:rules.bzl", "android_binary", "android_library")

alias(
    name = "android_example",
    actual = "//examples/android/com/java/bazel:hello_world",
)
