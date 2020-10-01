package afe

import (
	"testing"

	"github.com/spf13/afero"
)

// TODO use billy test suite
func Test(t *testing.T) {
	afs := afero.NewOsFs()
	fs := New(afs)
	filename := "testfile"
	newfilename := "newtestfile"

	_, err := fs.Create(filename)
	if err != nil {
		t.Error(err)
	}

	fi, err := fs.Stat(filename)
	if err != nil {
		t.Error(err)
	}
	if fi.Name() != filename {
		t.Errorf("wanted %s, got %s", filename, fi.Name())
	}

	err = fs.Rename(filename, newfilename)
	if err != nil {
		t.Error(err)
	}
	fi, err = fs.Stat(newfilename)
	if err != nil {
		t.Error(err)
	}
	if fi.Name() != newfilename {
		t.Errorf("wanted %s, got %s", newfilename, fi.Name())
	}

	err = fs.Remove(newfilename)
	if err != nil {
		t.Error(err)
	}
}
