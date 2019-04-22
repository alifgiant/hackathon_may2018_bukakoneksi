package workspace

import (
	"log"

	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/amenities"

	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi"
)

func CreateWorkspace(w bukakoneksi.Workspace) (bukakoneksi.Workspace, error) {
	w.AmenitiesTags = amenities.ConvertArrayToString(w.Amenities)
	if err := bukakoneksi.Transaction.Insert(&w); err != nil {
		log.Println(err)
		return w, err
	}

	var table bukakoneksi.Table
	err := bukakoneksi.Transaction.SelectOne(&table, "select * from tables where id = ?", w.TableID)
	if err != nil {
		log.Println(err)
		return w, err
	}

	table.EmptyWorkspace--
	_, err = bukakoneksi.Transaction.Update(&table)
	if err != nil {
		log.Println(err)
		return w, err
	}

	return w, nil
}
