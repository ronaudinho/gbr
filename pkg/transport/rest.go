package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-git/go-billy"
	"github.com/julienschmidt/httprouter"
)

// Transport
type Transport struct {
	bfs billy.Basic
}

// New creates new instance of Transport
func New(bfs billy.Basic) *Transport {
	return &Transport{
		bfs: bfs,
	}
}

// Create creates the named file with mode 0666 (before umask), truncating
// it if it already exists. If successful, methods on the returned File can
// be used for I/O; the associated file descriptor has mode O_RDWR.
func (t *Transport) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	filename := ps.ByName("filename")
	t.bfs.Create(filename)
}

// Open opens the named file for reading. If successful, methods on the
// returned file can be used for reading; the associated file descriptor has
// mode O_RDONLY.
func (t *Transport) Open(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("not yet implemented"))
}

// OpenFile is the generalized open call; most users will use Open or Create
// instead. It opens the named file with specified flag (O_RDONLY etc.) and
// perm, (0666 etc.) if applicable. If successful, methods on the returned
// File can be used for I/O.
func (t *Transport) OpenFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("not yet implemented"))
}

// Stat returns a FileInfo describing the named file.
func (t *Transport) Stat(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	filename := ps.ByName("filename")
	fi, err := t.bfs.Stat(filename)
	if err != nil {
		w.WriteHeader(http.StatusPaymentRequired)
		w.Write([]byte(fmt.Sprintf("stat: error: %s", err.Error())))
		return
	}

	b, _ := json.Marshal(struct {
		Name string `json:"name"`
		Size int64  `json:"size"`
	}{
		Name: fi.Name(),
		Size: fi.Size(),
	})
	w.Write(b)
	return
}

// Rename renames (moves) oldpath to newpath. If newpath already exists and
// is not a directory, Rename replaces it. OS-specific restrictions may
// apply when oldpath and newpath are in different directories.
func (t *Transport) Rename(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	oldpath := ps.ByName("filename")
	newpath := ps.ByName("newfilename")
	err := t.bfs.Rename(oldpath, newpath)
	if err != nil {
		w.WriteHeader(http.StatusPaymentRequired)
		w.Write([]byte(fmt.Sprintf("rename: error: %s", err.Error())))
		return
	}

	w.Write([]byte(fmt.Sprintf("rename: success: %s to %s\n", oldpath, newpath)))
	return
}

// Remove removes the named file or directory.
func (t *Transport) Remove(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	filename := ps.ByName("filename")
	err := t.bfs.Remove(filename)
	if err != nil {
		w.WriteHeader(http.StatusPaymentRequired)
		w.Write([]byte(fmt.Sprintf("remove: error: %s", err.Error())))
		return
	}

	w.Write([]byte(fmt.Sprintf("remove: success: %s\n", filename)))
	return
}

// Join joins any number of path elements into a single path, adding a
// Separator if necessary. Join calls filepath.Clean on the result; in
// particular, all empty strings are ignored. On Windows, the result is a
// UNC path if and only if the first path element is a UNC path.
func (t *Transport) Join(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("not yet implemented"))
}
