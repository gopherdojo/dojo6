package omikuji

import (
	"fmt"
	"net/http"
	"os"

	service "omikuji-app/pkg/api/domain/service/omikuji"
)

type omikujiHandler struct {
	omikujiService service.OmikujiService
}

func New(s service.OmikujiService) http.Handler {
	return &omikujiHandler{omikujiService: s}
}

func (h *omikujiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := h.omikujiService.Draw()
	if err != nil {
		fmt.Fprintf(os.Stderr, "抽選に失敗しました. err = %v\n", err)
	}
	fmt.Fprintf(w, "%s", res)
}
