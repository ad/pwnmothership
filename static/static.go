package static

import (
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/rakyll/statik/fs"
)

// FS stores link to statikFS if it exists
type FS struct {
	Statik   http.FileSystem
	StatikFS http.FileSystem
}

// FileReader describes methods of filesystem
type FileReader interface {
	ReadFile(path string) ([]byte, error)
}

// NewFS returns new FS instance, which will read from statik if it's available and from fs otherwise
func NewFS() *FS {
	f := &FS{}

	if statikFS, err := fs.New(); err == nil {
		log.Printf("[INFO] static files will be read from statik bindata")
		f.StatikFS = statikFS
	}

	if statikFS, err := fs.NewWithNamespace("static"); err == nil {
		log.Printf("[INFO] static will be read from zip file")
		f.Statik = statikFS
	}

	return f
}

// ReadFile depends on statik achieve exists
func (f *FS) ReadFile(path string) ([]byte, error) {
	if f.Statik != nil {
		return fs.ReadFile(f.Statik, filepath.Join("/", path))
	}
	if f.StatikFS != nil {
		return fs.ReadFile(f.StatikFS, filepath.Join("/", path))
	}
	return ioutil.ReadFile(filepath.Join("./", filepath.Clean(path)))
}
