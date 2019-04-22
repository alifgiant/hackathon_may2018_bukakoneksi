package workspace

import (
	"log"

	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi"
)

func DeleteWorkspace(ID int) error {
	var w bukakoneksi.Workspace

	err := bukakoneksi.Dbmap.SelectOne(&w, "select * from workspaces where id = ?", ID)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = bukakoneksi.Transaction.Delete(&w)
	if err != nil {
		log.Println(err)
		return err
	}

	var table bukakoneksi.Table
	err = bukakoneksi.Dbmap.SelectOne(&table, "select * from tables where id = ?", w.TableID)
	if err != nil {
		log.Println(err)
		return err
	}

	table.EmptyWorkspace++
	_, err = bukakoneksi.Transaction.Update(&table)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
