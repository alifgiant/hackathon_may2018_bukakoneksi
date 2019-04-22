package member

import (
	"log"

	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi"
)

func DeleteMember(ID string) error {
	var m bukakoneksi.Member
	err := bukakoneksi.Dbmap.SelectOne(&m, "select * from members where telegram_user_id = ?", ID)
	if err != nil {
		log.Println(err)
		return err
	}

	bukakoneksi.Transaction, err = bukakoneksi.Dbmap.Begin()
	if err != nil {
		bukakoneksi.Transaction.Rollback()
		log.Println(err)
		return err
	}

	_, err = bukakoneksi.Transaction.Exec("delete from members where telegram_user_id = ?", ID)
	if err != nil {
		bukakoneksi.Transaction.Rollback()
		log.Println(err)
		return err
	}

	_, err = bukakoneksi.Transaction.Exec("delete from workspaces where member_id = ?", m.ID)
	if err != nil {
		bukakoneksi.Transaction.Rollback()
		log.Println(err)
		return err
	}

	if err = bukakoneksi.Transaction.Commit(); err != nil {
		bukakoneksi.Transaction.Rollback()
		log.Println(err)
		return err
	}

	return nil
}
