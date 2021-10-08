package shortener

import (
	"errors"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const fileExt = ".url"

type OnDisk struct {
	basePath string
}

func NewOnDisk(path string) (*OnDisk, error) {
	err := os.MkdirAll(path, 0o700)
	return &OnDisk{path}, err
}

func (s *OnDisk) Get(id string) string {
	path := s.path(id)
	data, err := os.ReadFile(path)
	switch {
	case errors.Is(err, fs.ErrNotExist):
		return ""
	case err != nil:
		log.Println(err)
		return ""
	}
	return string(data)
}

func (s *OnDisk) Put(url string) string {
	id := generateID(url)
	f, err := os.OpenFile(s.path(id), os.O_CREATE|os.O_EXCL|os.O_RDWR, 0o600)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer f.Close()
	if _, err := io.Copy(f, strings.NewReader(url)); err != nil {
		panic(err)
	}
	return id
}

func (s *OnDisk) path(id string) string {
	return filepath.Join(s.basePath, id+fileExt)
}
