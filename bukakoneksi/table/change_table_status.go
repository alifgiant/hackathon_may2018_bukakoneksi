package table

import (
	"log"

	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi"
)

func ChangeTableStatus(id int, pointX float32, pointY float32) error {
	var table bukakoneksi.Table

	err := bukakoneksi.Dbmap.SelectOne(&table, "select * from tables where id = ?", id)
	if err != nil {
		log.Println(err)
		return err
	}

	table.PointX = pointX
	table.PointY = pointY
	table.Status = "map"
	_, err = bukakoneksi.Dbmap.Update(&table)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
