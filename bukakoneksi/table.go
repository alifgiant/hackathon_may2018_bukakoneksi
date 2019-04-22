package bukakoneksi

// Table struct
type Table struct {
	ID             int          `db:"id" json:"id"`
	OfficeFloorID  int          `db:"office_floor_id" json:"office_floor_id"`
	Name           string       `db:"name" json:"name"`
	Status         string       `db:"status" json:"status"` // status : map / unmap
	WorkspaceSize  int          `db:"workspace_size" json:"workspace_size"`
	EmptyWorkspace int          `db:"empty_workspace" json:"empty_workspace"`
	PointX         float32      `db:"point_x" json:"point_x"`
	PointY         float32      `db:"point_y" json:"point_y"`
	Workspaces     []Workspace  `db:"-" json:"workspaces"`
	OfficeFloor    *OfficeFloor `db:"-" json:"office_floor,omitempty"`
	Office         *Office      `db:"-" json:"office,omitempty"`
}
