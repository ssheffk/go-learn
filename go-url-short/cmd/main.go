package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-learn/go-url-short/internal/urlshort"
)

var (
	pathsFile = flag.String("pathsFile", "paths.json", "The file containing shortened paths to URL's")
)

func getFileBytes(fileName string) []byte {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Could not open files %s", fileName)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(f)
	if err != nil {
		log.Fatalf("Could not read file %s", fileName)
	}

	return buf.Bytes()
}

func main() {
	mux := defaultMux()
	flag.Parse()
	ext := filepath.Ext(*pathsFile)

	var handler http.Handler
	var err error

	if ext == ".yaml" {
		handler, err = urlshort.YAMLHandler(getFileBytes(*pathsFile), mux)
		if err != nil {
			panic(err)
		}
	} else if ext == ".json" {
		handler, err = urlshort.JSONHandler(getFileBytes(*pathsFile), mux)
		if err != nil {
			panic(err)
		}
	} else {
		log.Fatalf("Paths file need to be either a YAML or JSON")
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}
