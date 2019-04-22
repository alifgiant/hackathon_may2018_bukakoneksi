package city

import (
	"net/http"

	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/server/response"
	"github.com/julienschmidt/httprouter"
)

type Handler struct{}

func (h *Handler) Retrieve(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var cities []bukakoneksi.City

	cities, _ = RetrieveCities()

	response.OK(w, cities)
}
