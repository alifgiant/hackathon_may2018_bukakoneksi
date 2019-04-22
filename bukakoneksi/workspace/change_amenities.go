package workspace

import (
	"log"

	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/amenities"
)

func ChangeAmenities(ID int, amt []string) ([]string, error) {
	var workspace bukakoneksi.Workspace

	err := bukakoneksi.Dbmap.SelectOne(&workspace, "select * from workspaces where id = ?", ID)
	if err != nil {
		log.Println(err)
		return amt, err
	}

	workspace.AmenitiesTags = amenities.ConvertArrayToString(amt)
	workspace.Amenities = amt

	_, err = bukakoneksi.Dbmap.Update(&workspace)
	if err != nil {
		log.Println(err)
		return amt, err
	}

	return amt, nil
}
