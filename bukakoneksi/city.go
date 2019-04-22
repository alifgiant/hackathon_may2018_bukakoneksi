package bukakoneksi

// City struct
type City struct {
	ID   int    `db:"id" json:"-"`
	Name string `db:"name" json:"name"`
}
