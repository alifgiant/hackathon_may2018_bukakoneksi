package bukakoneksi

// Office struct
type Office struct {
	ID           int           `db:"id" json:"id"`
	Name         string        `db:"name" json:"name"`
	Address      string        `db:"address" json:"address"`
	City         string        `db:"city" json:"city"`
	LocationURL  string        `db:"location_url" json:"location_url"`
	OfficeFloors []OfficeFloor `db:"-" json:"office_floors"`
}
