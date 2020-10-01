package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ronaudinho/gbr/pkg/afe"
	"github.com/ronaudinho/gbr/pkg/transport"
	"github.com/spf13/afero"
)

func main() {
	afs := afero.NewOsFs()
	fs := afe.New(afs)
	trp := transport.New(fs)
	router := httprouter.New()
	router.GET("/:filename", trp.Stat)
	router.POST("/:filename", trp.Create)
	router.POST("/:filename/:newfilename", trp.Rename)
	router.DELETE("/:filename", trp.Remove)

	log.Fatal("serving :3195", http.ListenAndServe(":3195", router))
}
