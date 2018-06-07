http_archive(
    name = "build_bazel_rules_android",
    sha256 = "1e50a4227197edcd3712f1aeb92297aa3a1edcb0931464a8872a1c2fe77160ac",
    strip_prefix = "rules_android-0.1.0",
    urls = ["https://github.com/bazelbuild/rules_android/archive/v0.1.0.tar.gz"],
)

git_repository(
    name = "android_sdk_downloader",
    commit = "a08905c5571dc9a74027ec57c90ffad53d7f7efe",
    remote = "https://github.com/quittle/bazel_android_sdk_downloader",
)

load("@android_sdk_downloader//:rules.bzl", "android_sdk_repository")

android_sdk_repository(
    name = "androidsdk",
    api_level = 27,
    build_tools_version = "27.0.3",
    workspace_name = "com_example_workspace",
)

# Uncomment to use the Android NDK
# android_ndk_repository(
#     name = "androidndk",
# )
