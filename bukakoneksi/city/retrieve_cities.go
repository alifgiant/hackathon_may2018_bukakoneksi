package city

import "github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi"

func RetrieveCities() ([]bukakoneksi.City, error) {
	var cities []bukakoneksi.City

	_, err := bukakoneksi.Dbmap.Select(&cities, "select * from cities order by name asc")

	return cities, err
}
