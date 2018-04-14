# GoFish, The Package Manager

Features, usage and installation instructions can be found on the [homepage](https://gofi.sh).

## What does GoFish do?

GoFish is a cross-platform systems package manager, bringing the ease of use of Homebrew to
Linux and Windows.

```
$ gofish install go
==> Installing go...
🐠  go 1.10.1: installed in 2.307602197s
```

GoFish works across all three major operating systems (Windows, MacOS, and Linux). It installs
packages into its own directory and symlinks their files into /usr/local (or C:\ProgramData for Windows).
You can think of it as the cross-platform Homebrew.

## Want to add your project to the list of installable thingies?

Make a PR at [fishworks/fish-food](https://github.com/fishworks/fish-food)! Just make sure to follow the [Contributing Guide](https://gofi.sh#contributing) first.

## Troubleshooting

For troubleshooting, see the [Troubleshooting Guide](https://gofi.sh#troubleshooting).

## Security

Please email security issues to [Matt Fisher](mailto:matt.fisher+security-issues@fishworks.io).

## License

GoFish is licensed under the [Apache v2 License](LICENSE).
