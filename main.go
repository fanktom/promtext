package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

func main() {
	// Parse flags
	var port = flag.Int("port", 9100, "port for the server")
	var dir = flag.String("dir", "files", "folder to collect")
	flag.Parse()

	// Handle HTTP on /
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Read list of files in dir
		files, err := ioutil.ReadDir(*dir)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, file := range files {
			b, err := ioutil.ReadFile(path.Join(*dir, file.Name()))
			if err != nil {
				continue
			}
			fmt.Fprintf(w, string(b))
		}
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *port), nil))
}
