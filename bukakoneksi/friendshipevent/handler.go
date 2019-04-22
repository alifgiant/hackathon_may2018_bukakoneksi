package friendshipevent

import (
	"log"
	"net/http"

	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/server/response"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/utils"
	"github.com/julienschmidt/httprouter"
)

type Handler struct{}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var fe bukakoneksi.FriendshipEvent
	err := utils.Decode(r, &fe)
	if err != nil {
		log.Println(err)
		return
	}

	fe, err = CreateFriendshipEvent(fe)
	if err != nil {
		log.Println(err)
		return
	}

	response.OK(w, fe)
}

func (h *Handler) Retrieve(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	city := ""
	event := ""
	day := ""

	params := r.URL.Query()

	label := "city"
	if utils.IsQueryParamExist(params, label) {
		city = params[label][0]
	}

	label = "event"
	if utils.IsQueryParamExist(params, label) {
		event = params[label][0]
	}

	label = "day"
	if utils.IsQueryParamExist(params, label) {
		day = params[label][0]
	}

	res, err := RetrieveFriendshipEvent(city, event, day)
	if err != nil {
		log.Println(err)
		return
	}

	response.OK(w, res)
}
