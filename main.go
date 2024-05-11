package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const Port, Path, Name = ":8080", "/", "Cequallium"
const separator = "_____________________________________________________________"

var Footer = fmt.Sprintf("%v\n\n Probably best viewed in a terminal...\n\nCopyright (c) %v %v. All Rights Reserved.", separator, time.Now().Year(), Name)

func serve(w http.ResponseWriter, r *http.Request) {
	wd, _ := os.Getwd()
	file, err := os.ReadFile(fmt.Sprintf("%v/pages%v", wd, r.URL.Path))

	if err != nil {
		fmt.Fprintf(w, "404")
	} else {
		fmt.Fprintf(w, "[ %v ]\n\n", r.URL.Path)
		fmt.Fprintf(w, "%v", string(file))
		fmt.Fprint(w, Footer)
	}
}
func main() {
	fmt.Printf("whispering on localhost%v%v", Port, Path)
	http.HandleFunc(Path, serve)
	http.ListenAndServe(Port, nil)
}
