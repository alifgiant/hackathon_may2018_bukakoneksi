package bukakoneksi

// Amenities struct
type Amenities struct {
	ID   int    `db:"id" json:"-"`
	Name string `db:"name" json:"name"`
}
