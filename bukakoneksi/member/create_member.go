package member

import (
	"log"

	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi"
)

func CreateMember(m bukakoneksi.Member) (bukakoneksi.Member, error) {
	m.TelegramUsername = "@" + m.TelegramUsername
	if err := bukakoneksi.Dbmap.Insert(&m); err != nil {
		log.Println(err)
		return m, err
	}

	return m, nil
}
