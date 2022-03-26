package gofish

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/jdxcode/netrc"
	"github.com/mholt/archiver/v3"
	log "github.com/sirupsen/logrus"

	"github.com/fishworks/gofish/pkg/home"
	"github.com/fishworks/gofish/pkg/osutil"
	"github.com/fishworks/gofish/receipt"
	"github.com/fishworks/gofish/version"
)

// Food provides metadata to install a piece of software.
type Food struct {
	// The canonical name of the software.
	Name string
	// The repository where this food resides.
	Rig string
	// A (short) description of the software.
	Description string
	// The license identifier for the software.
	License string
	// The homepage URL for the software.
	Homepage string
	// Caveats inform the user about any Draft-specific caveats regarding this fish food.
	Caveats string
	// The version of the software.
	Version string
	// The list of binary distributions available for this fish food.
	Packages []*Package
	// The script to run before installation
	PreInstallScript string
	// The script to run after a successful installation
	PostInstallScript string
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

// Resource is a installable thingy that should be moved into /usr/local from the install path, such as an executable, manpages, libraries, etc.
type Resource struct {
	// Path is the path relative from the root of the unpacked archive to the resource. The resource is symlinked into the InstallPath and, if Executable is set, made executable (chmod +x).
	Path string
	// InstallPath is the destination path relative from /usr/local. The resource is symlinked from Path to the InstallPath and, if Executable is set, made executable (chmod +x).
	InstallPath string
	// Executable defines whether or not this resource should be made executable (chmod +x). This only applies for MacOS/Linux and can be ignored on Windows.
	Executable bool
}

// ErrCouldNotUnlink is returned when the 'unlink' operation does not succeed
type ErrCouldNotUnlink struct {
	Err error
}

func (e ErrCouldNotUnlink) Error() string {
	return fmt.Sprintf("could not unlink: %s", e.Err.Error())
}

// ErrCouldNotLink is returned when the 'link' operation does not succeed
type ErrCouldNotLink struct {
	Err error
}

func (e ErrCouldNotLink) Error() string {
	return fmt.Sprintf("could not link: %s", e.Err.Error())
}

// Install attempts to install the package, returning errors if it fails.
func (f *Food) Install() error {
	barrelDir := filepath.Join(home.Barrel(), f.Name, f.Version)
	pkg := f.GetPackage(runtime.GOOS, runtime.GOARCH)
	if pkg == nil {
		return fmt.Errorf("food '%s' does not support the current platform (%s/%s)", f.Name, runtime.GOOS, runtime.GOARCH)
	}
	u, err := url.Parse(pkg.URL)
	if err != nil {
		return fmt.Errorf("could not parse package URL '%s' as a URL: %v", pkg.URL, err)
	}
	cachedFilePath := filepath.Join(home.Cache(), fmt.Sprintf("%s-%s-%s-%s%s", f.Name, f.Version, pkg.OS, pkg.Arch, getExtension(u.Path)))
	if err := f.DownloadTo(pkg, cachedFilePath); err != nil {
		return err
	}
	if err := checksumVerifyPath(cachedFilePath, pkg.SHA256); err != nil {
		return fmt.Errorf("shasum verify check failed: %v", err)
	}

	if f.PreInstallScript != "" {
		cmd := exec.Command(f.PreInstallScript)
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	if err := os.MkdirAll(barrelDir, 0755); err != nil {
		return err
	}
	unarchiveOrCopy(cachedFilePath, barrelDir, u.Path)

	// This is just a safety check to make sure that there's nothing there when we link the package.
	err = f.Unlink(pkg)
	if err != nil {
		return ErrCouldNotUnlink{err}
	}

	// special case: gofish is replacing itself on windows
	// https://github.com/fishworks/gofish/issues/46
	if runtime.GOOS == "windows" && f.Name == "gofish" {
		gofishBinPath := filepath.Join(home.HomePrefix, "bin/gofish.exe")
		exists, err := osutil.Exists(gofishBinPath)
		if err != nil {
			return err
		}
		if exists {
			if err := os.Rename(gofishBinPath, fmt.Sprintf("%s.rotten", gofishBinPath)); err != nil {
				return err
			}
		}
	}
	if err := f.Link(pkg); err != nil {
		return ErrCouldNotLink{err}
	}

	if f.PostInstallScript != "" {
		cmd := exec.Command(f.PostInstallScript)
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	if f.Caveats != "" {
		fmt.Println(f.Caveats)
	}

	// write to the install receipt to record what's been done here
	receiptFilepath := filepath.Join(home.Barrel(), f.Name, receipt.ReceiptFilename)

	receiptFile, err := os.OpenFile(receiptFilepath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer receiptFile.Close()

	installReceipt, err := receipt.NewFromReader(receiptFile)
	if err != nil && err != io.EOF {
		return err
	}

	// rewind to the beginning to overwrite the receipt's contents
	receiptFile.Seek(0, 0)

	installReceipt.Name = f.Name
	installReceipt.Rig = f.Rig
	installReceipt.LastModified = time.Now()
	installReceipt.GoFishVersion = version.String()

	return installReceipt.Save(receiptFile)
}

// Uninstall attempts to uninstall the package, returning errors if it fails.
func (f *Food) Uninstall() error {
	pkg := f.GetPackage(runtime.GOOS, runtime.GOARCH)
	if pkg == nil {
		return nil
	}
	if f.Linked() {
		if err := f.Unlink(pkg); err != nil {
			return err
		}
	}
	barrelDir := filepath.Join(home.Barrel(), f.Name, f.Version)
	os.Remove(filepath.Join(home.Barrel(), f.Name, receipt.ReceiptFilename))
	return os.RemoveAll(barrelDir)
}

func unarchiveOrCopy(src, dest, urlPath string) error {

	// check and see if it can be unarchived by archiver
	if _, err := archiver.ByExtension(src); err == nil {
		return archiver.Unarchive(src, dest)
	}

	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(filepath.Join(dest, filepath.Base(urlPath)))
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
	barrelDir := filepath.Join(home.Barrel(), f.Name, f.Version)
	link, err := os.Readlink(filepath.Join(home.BinPath(), f.Name))
	if err != nil {
		return false
	}
	return strings.Contains(link, barrelDir)
}

// Link creates links to any linked resources owned by the package.
func (f *Food) Link(pkg *Package) error {
	barrelDir := filepath.Join(home.Barrel(), f.Name, f.Version)
	for _, r := range pkg.Resources {
		// TODO: run this in parallel
		destPath := filepath.Join(home.HomePrefix, r.InstallPath)
		if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil && !os.IsExist(err) {
			return err
		}
		if r.Executable {
			if err := os.Chmod(filepath.Join(barrelDir, r.Path), 0755); err != nil {
				return err
			}
		}
		if err := os.Symlink(filepath.Join(barrelDir, r.Path), destPath); err != nil {
			return err
		}
	}
	return nil
}

// Unlink removes any linked resources owned by the package.
func (f *Food) Unlink(pkg *Package) error {
	for _, r := range pkg.Resources {
		// TODO: check if the linked path we are about to remove is really owned by us
		if err := os.RemoveAll(filepath.Join(home.HomePrefix, r.InstallPath)); err != nil {
			return err
		}
	}
	return nil
}

// Lint analyses a given fish food for potential errors, returning a list of errors.
func (f *Food) Lint() (errs []error) {
	var wg sync.WaitGroup
	for _, pkg := range f.Packages {
		wg.Add(1)
		go func(pkg *Package) {
			defer wg.Done()
			u, err := url.Parse(pkg.URL)
			if err != nil {
				errs = append(errs, fmt.Errorf("could not parse package URL '%s' as a URL: %v", pkg.URL, err))
			}
			cachedFilePath := filepath.Join(home.Cache(), fmt.Sprintf("%s-%s-%s-%s%s", f.Name, f.Version, pkg.OS, pkg.Arch, getExtension(u.Path)))
			if err := f.DownloadTo(pkg, cachedFilePath); err != nil {
				errs = append(errs, err)
			}
			if err := checksumVerifyPath(cachedFilePath, pkg.SHA256); err != nil {
				errs = append(errs, fmt.Errorf("shasum verify check failed: %v", err))
			}
			if err := installVerify(f, pkg, cachedFilePath); err != nil {
				errs = append(errs, fmt.Errorf("install verify check failed: %v", err))
			}
		}(pkg)
	}
	wg.Wait()
	return
}

func installVerify(f *Food, pkg *Package, src string) error {
	barrel := filepath.Join(home.Cache(), "barrel")
	barrelDir := filepath.Join(barrel, f.Name, f.Version, pkg.OS, pkg.Arch)

	u, err := url.Parse(pkg.URL)
	if err != nil {
		return fmt.Errorf("could not parse package URL '%s' as a URL: %v", pkg.URL, err)
	}

	err = os.RemoveAll(barrelDir)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(barrelDir, 0755); err != nil && !os.IsExist(err) {
		return err
	}

	err = unarchiveOrCopy(src, barrelDir, u.Path)
	if err != nil {
		return err
	}

	for _, r := range pkg.Resources {
		rPath := r.Path
		// If we are validating windows packages on a non-windows machine
		// we need to change the directory separator
		if runtime.GOOS != "windows" {
			rPath = strings.ReplaceAll(r.Path, "\\", "/")
		}
		resourcePath := filepath.Join(barrelDir, rPath)
		_, err := os.Stat(resourcePath)
		if err != nil {
			return err
		}
	}
	return nil
}

func getExtension(path string) string {
	urlParts := strings.Split(path, "/")
	parts := strings.Split(urlParts[len(urlParts)-1], ".")
	if len(parts) < 2 {
		return filepath.Ext(path)
	}
	return "." + strings.Join([]string{parts[len(parts)-2], parts[len(parts)-1]}, ".")
}

// DownloadTo downloads a particular package to filePath, returning any errors if encountered.
func (f *Food) DownloadTo(pkg *Package, filePath string) error {
	var success = true
	if err := downloadCachedFileToPath(filePath, pkg.URL); err != nil {
		success = false
		log.Errorln(err)
		// try using the mirrors
		for i := range pkg.Mirrors {
			if err := downloadCachedFileToPath(filePath, pkg.Mirrors[i]); err == nil {
				success = true
				break
			} else {
				log.Errorln(err)
			}
		}
	}
	if !success {
		return fmt.Errorf("failed to download package for OS/arch %s/%s with URL %s to filepath %s", pkg.OS, pkg.Arch, pkg.URL, filePath)
	}
	return nil
}

// userNetrcCredentials will attempt to load login and password from the users netrc file
func userNetrcCredentials(host string) (login, password string) {
	var netrcMachine *netrc.Machine
	for _, netrcFilePath := range []string{os.Getenv("NETRC"), home.GPGNetrc(), home.Netrc()} {
		if netrcFile, err := netrc.Parse(netrcFilePath); err != nil {
			if !os.IsNotExist(err) {
				log.Errorln(err)
			}
		} else {
			netrcMachine = netrcFile.Machine(host)
			return netrcMachine.Get("login"), netrcMachine.Get("password")
		}
	}
	return "", ""
}

// addCredentialsToRequest will attempt to attach relevant credentials to the http.Request
func addCredentialsToRequest(req *http.Request) {
	login, password := userNetrcCredentials(req.Host)
	if login != "" && password != "" {
		req.SetBasicAuth(login, password)
	}
}

// downloadCachedFileToPath will download a file from the given url to a directory, returning the
// path to the cached file. If it already exists, it'll skip downloading the file and just return
// the path to the cached file.
func downloadCachedFileToPath(filePath string, url string) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/octet-stream")

	addCredentialsToRequest(req)

	if _, err = os.Stat(filePath); err == nil {
		return nil
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func checksumVerifyPath(path string, checksum string) error {
	hasher := sha256.New()
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := io.Copy(hasher, f); err != nil {
		return err
	}

	actualChecksum := fmt.Sprintf("%x", hasher.Sum(nil))
	if strings.Compare(actualChecksum, strings.ToLower(checksum)) != 0 {
		return fmt.Errorf("checksums differ for %s: expected '%s', got '%s'", path, checksum, actualChecksum)
	}
	return nil
}
