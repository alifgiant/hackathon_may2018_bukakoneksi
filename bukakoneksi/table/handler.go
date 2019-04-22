package table

import (
	"log"
	"net/http"

	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/server/response"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/utils"
	"github.com/julienschmidt/httprouter"
)

type Handler struct{}

func (h *Handler) Retrieve(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var tables []bukakoneksi.Table

	criteria := make(map[string]string)
	criteria["office_floor_id"] = ""
	// label search
	criteria["name"] = ""
	criteria["empty_workspace"] = ""
	criteria["member_telegram"] = ""
	criteria["source"] = ""

	params := r.URL.Query()

	label := "office_floor_id"
	if utils.IsQueryParamExist(params, label) {
		criteria["office_floor_id"] = params[label][0]
	}

	label = "source"
	if utils.IsQueryParamExist(params, label) {
		criteria["source"] = params[label][0]
	}

	label = "search"
	search := ""
	if utils.IsQueryParamExist(params, label) {
		search = params[label][0]
	}
	if criteria["office_floor_id"] != "" && search == "empty" {
		criteria["empty_workspace"] = "empty"
	} else if search != "" && search[0] == '@' {
		criteria["member_telegram"] = search
	} else if search != "" {
		criteria["name"] = search
	}

	tables, _ = RetrieveTableByCriteria(criteria)

	response.OK(w, tables)
}

func (h *Handler) ChangeStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var status struct {
		ID     int     `json:"id"`
		PointX float32 `json:"point_x"`
		PointY float32 `json:"point_y"`
	}

	err := utils.Decode(r, &status)
	if err != nil {
		log.Println(err)
		return
	}

	err = ChangeTableStatus(status.ID, status.PointX, status.PointY)
	if err != nil {
		log.Println(err)
		return
	}

	response.OK(w, status)
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

	table, err := DeleteTable(id.ID)
	if err != nil {
		log.Println(err)
		return
	}

	response.OK(w, table)
	return
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var t bukakoneksi.Table

	err := utils.Decode(r, &t)
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

	table, err := CreateTable(t)
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

	response.OK(w, table)
	return
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var t bukakoneksi.Table

	err := utils.Decode(r, &t)
	if err != nil {
		log.Println(err)
		return
	}

	tableEdit, err := TableExist(t)
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

	updated, err := UpdateTable(t, tableEdit)
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

	response.OK(w, updated)
	return
}
