package table

import (
	"log"

	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi"
)

func DeleteTable(id int) (bukakoneksi.Table, error) {
	var table bukakoneksi.Table

	err := bukakoneksi.Dbmap.SelectOne(&table, "select * from tables where id = ?", id)
	if err != nil {
		log.Println(err)
		return table, err
	}

	bukakoneksi.Transaction, err = bukakoneksi.Dbmap.Begin()
	if err != nil {
		bukakoneksi.Transaction.Rollback()
		log.Println(err)
		return table, err
	}

	_, err = bukakoneksi.Transaction.Delete(&table)
	if err != nil {
		bukakoneksi.Transaction.Rollback()
		log.Println(err)
		return table, err
	}

	_, err = bukakoneksi.Transaction.Exec("delete from workspaces where table_id = ?", table.ID)
	if err != nil {
		bukakoneksi.Transaction.Rollback()
		log.Println(err)
		return table, err
	}

	err = bukakoneksi.Transaction.Commit()
	if err != nil {
		bukakoneksi.Transaction.Rollback()
		log.Println(err)
		return table, err
	}

	return table, nil
}
