package handlers

import (
	"net/http"

	"github.com/nayan9229/go-load-generator/views/auth"
)

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, auth.Login())
}
