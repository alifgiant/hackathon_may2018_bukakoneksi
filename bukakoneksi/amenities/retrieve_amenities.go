package amenities

import (
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi"
)

func RetrieveAmenities() ([]bukakoneksi.Amenities, error) {
	var amenities []bukakoneksi.Amenities

	_, err := bukakoneksi.Dbmap.Select(&amenities, "select * from amenities order by name asc")

	return amenities, err
}
