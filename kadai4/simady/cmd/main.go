package main

import (
	"omikuji-app/pkg/api"
)

func main() {
	api.Serve(":8081")
}
