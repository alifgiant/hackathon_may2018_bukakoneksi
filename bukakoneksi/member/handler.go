package member

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
	member, err := RetrieveMember()
	if err != nil {
		log.Println(err)
		return
	}

	response.OK(w, member)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var member bukakoneksi.Member
	err := utils.Decode(r, &member)
	if err != nil {
		log.Println(err)
		return
	}

	err = DeleteMember(member.TelegramUserID)
	if err != nil {
		log.Println(err)
		return
	}

	response.OK(w, member)
	return
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var member bukakoneksi.Member
	err := utils.Decode(r, &member)
	if err != nil {
		log.Println(err)
		return
	}

	m, err := CreateMember(member)
	if err != nil {
		log.Println(err)
		return
	}

	mPic, err := RetrieveTelegramPicture(m.TelegramUserID)
	if err != nil {
		log.Println(err)
		m.PictureURL = utils.GetPic("default.jpg")
		m.PictureImage = "default.jpg"
	} else {
		err = utils.SavePic(mPic, m.TelegramUserID+".jpg")
		m.PictureURL = utils.GetPic(m.TelegramUserID + ".jpg")
		m.PictureImage = m.TelegramUserID + ".jpg"
	}

	_, err = bukakoneksi.Dbmap.Update(&m)
	if err != nil {
		log.Println(err)
		return
	}

	response.OK(w, m)
	return
}
