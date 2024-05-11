package main

import (
	"fmt"
	"net/http"
	"os"
)

const Port, Path = ":8080", "/"

func serve(w http.ResponseWriter, r *http.Request) {
	wd, _ := os.Getwd()
	file, err := os.ReadFile(fmt.Sprintf("%v/pages%v", wd, r.URL.Path))

	if err != nil {
		fmt.Fprintf(w,"404")
	} else {
		fmt.Fprintf(w, string(file))
	}
}
func main() {
	http.HandleFunc(Path, serve)
	http.ListenAndServe(Port, nil)
}
