# Fish, The Package Manager

Features, usage and installation instructions can be found in the [documentation](docs/README.md).

## What does Fish do?

Fish is a cross-platform systems package manager, bringing the ease of use of Homebrew into
Linux and Windows.

```
$ fish install helm
==> Installing helm...
🐠  helm 2.8.2: installed in 1.462258159s
```

Fish works across all three major operating systems (Windows, MacOS, and Linux). It installs
packages into its own directory and symlinks their files into /usr/local (or C:\ for Windows).
You can think of it as the cross-platform Homebrew.

Fish takes the ideas of [Homebrew Formulas][formula] to the next level by installing pre-packaged
tools. Fish food are simple Lua scripts:

```lua
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
            binpath = "windows-amd64/draft"
        }
    }
}
```

## Troubleshooting

For troubleshooting, see the [Toubleshooting Guide](docs/troubleshooting.md).

## Security

Please email security issues to <mailto:matt.fisher+security-issues@fishworks.io>.

## License

Fish is licensed under the [Apache v2 License](LICENSE).

[formula]: https://docs.brew.sh/Formula-Cookbook#homebrew-terminology
