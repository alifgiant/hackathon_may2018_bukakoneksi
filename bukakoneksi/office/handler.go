package office

import (
	"log"
	"net/http"

	"github.com/bukalapak/annex/server/response"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/utils"
	"github.com/julienschmidt/httprouter"
)

type Handler struct{}

func (h *Handler) Retrieve(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var offices []bukakoneksi.Office

	params := r.URL.Query()

	label := "city"
	city := ""
	if utils.IsQueryParamExist(params, label) {
		city = params[label][0]
	}

	offices, _ = RetrieveOffice(city)

	response.OK(w, offices)
}

func (h *Handler) Maps(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	offices, err := GetAll()
	if err != nil {
		log.Println(err)
		return
	}

	response.OK(w, offices)
	return
}
