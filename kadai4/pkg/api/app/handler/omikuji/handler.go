package omikuji

import (
	"bytes"
	"encoding/json"
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
	ctx := r.Context()
	rs := h.omikujiService.Draw(ctx)
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(rs); err != nil {
		fmt.Fprintf(os.Stderr, "抽選に失敗しました. err = %v\n", err)
		http.Error(w, "抽選に失敗しました.", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s", buf.String())
}
