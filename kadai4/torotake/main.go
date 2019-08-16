package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gopherdojo/dojo6/kadai4/torotake/pkg/omikuji"
)

var port int

func init() {
	// -p=[ポート番号] default : 8080
	flag.IntVar(&port, "p", 8080, "listen port number")
	flag.Parse()
}

func main() {
	server := omikuji.Server{}
	http.HandleFunc("/", server.Handler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}
