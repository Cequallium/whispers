package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// files will be served in ./pages/
const Port, UrlPath, Name, PagePath = ":8080", "/", "Author", "pages"

var separator string = strings.Repeat("_", 60)

var Footer = fmt.Sprintf(
	"%v\n\n Probably best viewed in a terminal...\n\nCopyright (c) %v %v. All Rights Reserved.\n",
	separator,
	time.Now().Year(),
	Name,
)

func serve(w http.ResponseWriter, r *http.Request) {
	wd, _ := os.Getwd()
	file, err := os.ReadFile(fmt.Sprintf("%v/%v%v", wd, PagePath, r.URL.Path))

	if err != nil {
		fmt.Fprintf(w, "404\n")
	} else {
		fmt.Fprintf(w, "[ %v ]\n\n", strings.TrimPrefix(r.URL.Path, "/"))
		fmt.Fprint(w, string(file))
		fmt.Fprintf(w, "%v\n", Footer)
	}
}

func main() {
	fmt.Printf("whispering on http://localhost%v%v\n", Port, UrlPath)
	http.HandleFunc(UrlPath, serve)
	http.ListenAndServe(Port, nil)
}
