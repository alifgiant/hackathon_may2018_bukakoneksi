package workspace

import (
	"log"
	"net/http"

	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/server/response"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/utils"
	"github.com/julienschmidt/httprouter"
)

type Handler struct{}

func (h *Handler) Assign(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var workspace bukakoneksi.Workspace

	err := utils.Decode(r, &workspace)
	if err != nil {
		log.Println(err)
		return
	}

	bukakoneksi.Transaction, err = bukakoneksi.Dbmap.Begin()
	if err != nil {
		bukakoneksi.Transaction.Rollback()
		log.Println(err)
		return
	}

	work, err := CreateWorkspace(workspace)
	if err != nil {
		bukakoneksi.Transaction.Rollback()
		log.Println(err)
		return
	}

	err = bukakoneksi.Transaction.Commit()
	if err != nil {
		bukakoneksi.Transaction.Rollback()
		log.Println(err)
		return
	}

	response.OK(w, work)
	return
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var id struct {
		ID int `json:"id"`
	}

	err := utils.Decode(r, &id)
	if err != nil {
		log.Println(err)
		return
	}

	bukakoneksi.Transaction, err = bukakoneksi.Dbmap.Begin()
	if err != nil {
		bukakoneksi.Transaction.Rollback()
		log.Println(err)
		return
	}

	err = DeleteWorkspace(id.ID)
	if err != nil {
		bukakoneksi.Transaction.Rollback()
		log.Println(err)
		return
	}

	err = bukakoneksi.Transaction.Commit()
	if err != nil {
		bukakoneksi.Transaction.Rollback()
		log.Println(err)
		return
	}

	response.OK(w, id)
	return
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var data struct {
		ID        int      `json:"id"`
		Amenities []string `json:"amenities"`
	}

	err := utils.Decode(r, &data)
	if err != nil {
		log.Println(err)
		return
	}

	amenities, err := ChangeAmenities(data.ID, data.Amenities)
	if err != nil {
		log.Println(err)
		return
	}

	response.OK(w, amenities)
}
