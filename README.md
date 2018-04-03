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
local name = "helm"
local version = "2.8.2"

food = {
    name = name,
    description = "The Kubernetes Package Manager",
    homepage = "https://github.com/kubernetes/helm",
    version = version,
    packages = {
        {
            os = "darwin",
            arch = "amd64",
            url = "https://storage.googleapis.com/kubernetes-helm/helm-v" .. version .. "-darwin-amd64.tar.gz",
            sha256 = "a0a8cf462080b2bc391f38b7cf617618b189cdef9f071c06fa0068c2418cc413",
            binpath = "darwin-amd64/" .. name
        },
        {
            os = "linux",
            arch = "amd64",
            url = "https://storage.googleapis.com/kubernetes-helm/helm-v" .. version .. "-linux-amd64.tar.gz",
            sha256 = "614b5ac79de4336b37c9b26d528c6f2b94ee6ccacb94b0f4b8d9583a8dd122d3",
            binpath = "linux-amd64/" .. name
        },
        {
            os = "windows",
            arch = "amd64",
            url = "https://storage.googleapis.com/kubernetes-helm/helm-v" .. version .. "-windows-amd64.tar.gz",
            sha256 = "cb6ea5d60f202c752f1f0777e4bebd98c619a2c18e52468df7a302e783216f23",
            binpath = "windows-amd64\\" .. name
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
