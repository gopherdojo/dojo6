package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gopherdojo/dojo6/kadai4/pei/pkg/fortune"
)

var port int

func init() {
	flag.IntVar(&port, "p", 8080, "server port")	
}

func main() {
	flag.Parse()
	f := fortune.NewFortune(fortune.Clock{})
	http.HandleFunc("/", f.Handler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
