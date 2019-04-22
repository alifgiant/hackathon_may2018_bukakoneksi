package bukakoneksi

// Workspace struct
type Workspace struct {
	ID            int      `db:"id" json:"id"`
	TableID       int      `db:"table_id" json:"table_id"`
	MemberID      int      `db:"member_id" json:"member_id"`
	MemberName    string   `db:"-" json:"member_name"`
	Member        Member   `db:"-" json:"member"`
	Position      int      `db:"position" json:"position"`
	AmenitiesTags string   `db:"amenities_tags" json:"-"`
	Amenities     []string `db:"-" json:"amenities"`
}
