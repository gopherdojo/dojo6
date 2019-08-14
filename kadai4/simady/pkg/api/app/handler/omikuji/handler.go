package omikuji

import (
	"fmt"
	"net/http"
	"os"

	interactor "omikuji-app/pkg/api/app/interactor/omikuji"
)

type omikujiHandler struct {
	omikujiInteractor interactor.OmikujiInteractor
}

func New(i interactor.OmikujiInteractor) http.Handler {
	return &omikujiHandler{omikujiInteractor: i}
}

func (h *omikujiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := h.omikujiInteractor.Draw(r.Context())
	if err != nil {
		fmt.Fprintf(os.Stderr, "抽選に失敗しました. err = %v\n", err)
		http.Error(w, "抽選に失敗しました.", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s", res)
}
