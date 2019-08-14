package api

import (
	"fmt"
	"net/http"

	"os"

	omikujiHandler "omikuji-app/pkg/api/app/handler/omikuji"
	omikujiInteractor "omikuji-app/pkg/api/app/interactor/omikuji"
	"omikuji-app/pkg/api/app/middleware"
	"omikuji-app/pkg/api/app/presenter"
	omikujiService "omikuji-app/pkg/api/domain/service/omikuji"
)

func Serve(addr string) {
	jsonPresenter := presenter.New()
	service := omikujiService.New()
	interactor := omikujiInteractor.New(jsonPresenter, service)
	handler := omikujiHandler.New(interactor)
	http.Handle("/", middleware.With(handler, middleware.ContextMiddleWare{}, middleware.ResponseHeaderMiddleWare{}))
	http.Handle("/favicon.ico", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Favicon is not set.", http.StatusNotFound)
	}))
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to serve. err = %v\n", err)
	}
}
