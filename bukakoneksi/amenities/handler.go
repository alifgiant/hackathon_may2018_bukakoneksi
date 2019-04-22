package amenities

import (
	"net/http"

	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/server/response"
	"github.com/julienschmidt/httprouter"
)

type Handler struct{}

func (h *Handler) Retrieve(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var amenities []bukakoneksi.Amenities

	amenities, _ = RetrieveAmenities()

	response.OK(w, amenities)
}
