package api

import (
	"fmt"
	"net/http"
	"os"

	"omikuji-app/pkg/api/app/middleware"

	omikujiHandler "omikuji-app/pkg/api/app/handler/omikuji"
	omikujiService "omikuji-app/pkg/api/domain/service/omikuji"
)

func Serve(addr string) {
	http.Handle("/", middleware.With(omikujiHandler.New(omikujiService.New()), middleware.ContextMiddleWare{}, middleware.ResponseHeaderMiddleWare{}))
	http.Handle("/favicon.ico", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Favicon is not set.", http.StatusNotFound)
	}))
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to serve. err = %v\n", err)
	}
}
