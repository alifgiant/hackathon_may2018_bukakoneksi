package office

import (
	"log"

	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/utils"
)

func RetrieveOffice(city string) ([]bukakoneksi.Office, error) {
	var offices []bukakoneksi.Office

	_, err := bukakoneksi.Dbmap.Select(&offices, "select * from offices where city = ? order by name asc", city)

	for i := range offices {
		var officeFloors []bukakoneksi.OfficeFloor

		_, err = bukakoneksi.Dbmap.Select(&officeFloors, "select * from office_floors where office_id = ?", offices[i].ID)

		for j := range officeFloors {
			if officeFloors[j].FloorImage != "" {
				officeFloors[j].FloorURL = utils.GetPic(officeFloors[j].FloorImage)
			}
		}

		offices[i].OfficeFloors = officeFloors
	}

	return offices, err
}

func GetAll() ([]bukakoneksi.Office, error) {
	var offices []bukakoneksi.Office

	_, err := bukakoneksi.Dbmap.Select(&offices, "select * from offices")
	if err != nil {
		log.Println(err)
		return offices, err
	}

	return offices, nil
}
