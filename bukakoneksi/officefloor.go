package bukakoneksi

// OfficeFloor struct
type OfficeFloor struct {
	ID         int    `db:"id" json:"id"`
	OfficeID   int    `db:"office_id" json:"-"`
	Name       string `db:"name" json:"name"`
	FloorImage string `db:"floor_image" json:"-"`
	FloorURL   string `db:"-" json:"floor_url"`
}
