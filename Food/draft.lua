local version = "0.12.0"

food = {
    name = "draft",
    description = "A tool for developers to create cloud-native applications on Kubernetes",
    homepage = "https://github.com/Azure/draft",
    version = version,
    packages = {
        {
            os = "darwin",
            arch = "amd64",
            url = "https://azuredraft.blob.core.windows.net/draft/draft-v" .. version .. "-darwin-amd64.tar.gz",
            sha256 = "5caa5cc89d81f193615e3ad55f2c08be59052c3309f7c37d0ed0136d54b82228",
            binpath = "darwin-amd64/draft"
        },
        {
            os = "linux",
            arch = "amd64",
            url = "https://azuredraft.blob.core.windows.net/draft/draft-v" .. version .. "-linux-amd64.tar.gz",
            sha256 = "89db5727cab7e0f295de149d914eaf32f4adcecabbb030a03300fca58be85b37",
            binpath = "linux-amd64/draft"
        },
        {
            os = "windows",
            arch = "amd64",
            url = "https://azuredraft.blob.core.windows.net/draft/draft-v" .. version .. "-windows-amd64.tar.gz",
            sha256 = "ec1a5f5054e6c1a00477a1a2291155039165630891a5b10ed2ebb8947db16480",
            binpath = "windows-amd64\\draft"
        }
    },
    caveats = [[Be aware that Draft is currently experimental and does not have a stable release yet. We make no backwards-compatible guarantees between releases.

When upgrading, make sure to `rm -rf ~/.draft` before bootstrapping Draft according to the installation guide:

    https://github.com/Azure/draft/blob/v]] .. version .. [[/docs/install.md

If you bootstrapped an application using `draft create`, you'll also want to remove the files `draft create` generated before running `draft create && draft up` again.

Please make sure to read the release notes for further information:

    https://github.com/Azure/draft/releases/tag/v]] .. version .. [[]]
}
