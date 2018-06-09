http_archive(
    name = "build_bazel_rules_android",
    sha256 = "1e50a4227197edcd3712f1aeb92297aa3a1edcb0931464a8872a1c2fe77160ac",
    strip_prefix = "rules_android-0.1.0",
    urls = ["https://github.com/bazelbuild/rules_android/archive/v0.1.0.tar.gz"],
)

android_sdk_repository(
    name = "androidsdk",
)

# Uncomment to use the Android NDK
android_ndk_repository(
    name = "androidndk",
)
