package afe

import (
	"os"
	"path/filepath"

	"github.com/go-git/go-billy"
	"github.com/spf13/afero"
)

const defaultDirectoryMode = 0755

var (
	_ billy.Basic = new(FS)
)

// FS wraps afero.Fs
type FS struct {
	a    afero.Fs
	path string
}

// New creates a new FS
func New(fs afero.Fs) *FS {
	return NewPath(fs, "/")
}

// NewPath creates a new FS with the given path
func NewPath(fs afero.Fs, path string) *FS {
	return &FS{
		a:    fs,
		path: path,
	}
}

// Create implements billy.Basic interface
func (f *FS) Create(filename string) (billy.File, error) {
	file, err := f.a.Create(filename)
	if err != nil {
		return nil, err
	}

	return NewFile(filename, file), nil
}

// Open implements billy.Basic interface
func (f *FS) Open(filename string) (billy.File, error) {
	file, err := f.a.Open(filename)
	if err != nil {
		return nil, err
	}

	return NewFile(filename, file), nil
}

// OpenFile implements billy.Basic interface
func (f *FS) OpenFile(filename string, flag int, perm os.FileMode) (billy.File, error) {
	file, err := f.a.OpenFile(filename, flag, perm)
	if err != nil {
		return nil, err
	}

	return NewFile(filename, file), nil
}

// Stat implements billy.Basic interface
func (f *FS) Stat(filename string) (os.FileInfo, error) {
	return f.a.Stat(filename)
}

// Rename implements billy.Basic interface
func (f *FS) Rename(oldpath string, newpath string) error {
	return f.a.Rename(oldpath, newpath)
}

// Remove implements billy.Basic interface
func (f *FS) Remove(filename string) error {
	return f.a.Remove(filename)
}

// Join implements billy.Basic interface
func (f *FS) Join(elem ...string) string {
	return filepath.Join(elem...)
}
