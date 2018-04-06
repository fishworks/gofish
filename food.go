package fish

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/docker/docker/pkg/archive"

	"github.com/fishworks/fish/pkg/osutil"
)

// Food provides metadata to install a piece of software.
type Food struct {
	// The canonical name of the software.
	Name string
	// A (short) description of the software.
	Description string
	// The homepage URL for the software.
	Homepage string
	// Caveats inform the user about any Draft-specific caveats regarding this fish food.
	Caveats string
	// The version of the software.
	Version string
	// The list of binary distributions available for this fish food.
	Packages []*Package
}

// Package provides metadata to install a piece of software on a given operating system and architecture.
type Package struct {
	// the running program's operating system target. One of darwin, linux, windows, and so on.
	OS string
	// the running program's architecture target. One of 386, amd64, arm, s390x, and so on.
	Arch      string
	Resources []*Resource
	// The URL used to download the binary distribution for this version of the fish food. The file must be a gzipped tarball (.tar.gz) or a zipfile (.zip) for unpacking.
	URL string
	// Additional URLs for this version of the fish food.
	Mirrors []string
	// To verify the cached download's integrity and security, we verify the SHA-256 hash matches what we've declared in the fish food.
	SHA256 string
}

type Resource struct {
	// Path is the path relative from the root of the unpacked archive to the resource. The resource is symlinked into the InstallPath and, if Executable is set, made executable (chmod +x).
	Path string
	// InstallPath is the destination path relative from /usr/local. The resource is symlinked from Path to the InstallPath and, if Executable is set, made executable (chmod +x).
	InstallPath string
	// Executable defines whether or not this resource should be made executable (chmod +x). This only applies for MacOS/Linux and can be ignored on Windows.
	Executable bool
}

// Install attempts to install the package, returning errors if it fails.
func (f *Food) Install() error {
	pkg := f.GetPackage(runtime.GOOS, runtime.GOARCH)
	if pkg == nil {
		return fmt.Errorf("food '%s' does not support the current platform (%s/%s)", f.Name, runtime.GOOS, runtime.GOARCH)
	}
	cachedFilePath, err := downloadCachedFileToPath(UserHome(UserHomePath).Cache(), pkg.URL)
	if err != nil {
		return err
	}
	barrelDir := filepath.Join(Home(HomePath).Barrel(), f.Name, f.Version)

	if err := os.MkdirAll(barrelDir, 0755); err != nil {
		return err
	}

	unarchiveOrCopy(cachedFilePath, barrelDir)
	f.Unlink(pkg)
	if err := f.Link(pkg); err != nil {
		return err
	}
	// This is just a safety check to make sure that there's nothing there when we link the package.
	if f.Caveats != "" {
		fmt.Println(f.Caveats)
	}
	return nil
}

// Installed checks to see if this fish food is installed. This is actually just a check for if the
// directory exists and is not empty.
func (f *Food) Installed() bool {
	files, err := ioutil.ReadDir(filepath.Join(Home(HomePath).Barrel(), f.Name, f.Version))
	if err != nil {
		return false
	}
	return len(files) > 0
}

// Uninstall attempts to uninstall the package, returning errors if it fails.
func (f *Food) Uninstall() error {
	pkg := f.GetPackage(runtime.GOOS, runtime.GOARCH)
	if pkg == nil {
		return nil
	}
	if err := f.Unlink(pkg); err != nil {
		return err
	}
	barrelDir := filepath.Join(Home(HomePath).Barrel(), f.Name, f.Version)
	return os.RemoveAll(barrelDir)
}

func unarchiveOrCopy(src, dest string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	if archive.IsArchivePath(src) {
		return archive.Untar(in, dest, &archive.TarOptions{NoLchown: true})
	}
	out, err := os.Create(filepath.Join(dest, filepath.Base(src)))
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	return err
}

// GetPackage does a lookup for a package supporting the given os/arch. If none were found, this
// returns nil.
func (f *Food) GetPackage(os, arch string) *Package {
	for _, pkg := range f.Packages {
		if pkg.OS == os && pkg.Arch == arch {
			return pkg
		}
	}
	return nil
}

// Linked checks to see if a particular package owned by this fish food is linked to /usr/local/bin.
// This is just a check if the binaries symlinked in /usr/local/bin link back to the barrel.
func (f *Food) Linked() bool {
	barrelDir := filepath.Join(Home(HomePath).Barrel(), f.Name, f.Version)
	link, err := os.Readlink(filepath.Join(BinPath, f.Name))
	if err != nil {
		return false
	}
	return strings.Contains(link, barrelDir)
}

func (f *Food) Link(pkg *Package) error {
	barrelDir := filepath.Join(Home(HomePath).Barrel(), f.Name, f.Version)
	for _, r := range pkg.Resources {
		// TODO: run this in parallel
		destPath := filepath.Join(HomePrefix, r.InstallPath)
		if r.Executable {
			if err := os.Chmod(filepath.Join(barrelDir, r.Path), 0755); err != nil {
				return err
			}
		}
		if err := osutil.SymlinkWithFallback(filepath.Join(barrelDir, r.Path), destPath); err != nil {
			return err
		}
	}
	return nil
}

func (f *Food) Unlink(pkg *Package) error {
	for _, r := range pkg.Resources {
		if err := os.RemoveAll(filepath.Join(HomePrefix, r.InstallPath)); err != nil {
			return err
		}
	}
	return nil
}

// downloadCachedFileToPath will download a file from the given url to a directory, returning the
// path to the cached file. If it already exists, it'll skip downloading the file and just return
// the path to the cached file.
func downloadCachedFileToPath(dir string, url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	filePath := filepath.Join(dir, path.Base(req.URL.Path))

	if _, err = os.Stat(filePath); err == nil {
		return filePath, nil
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return filePath, err
}
