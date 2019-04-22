package table

import (
	"fmt"
	"strings"

	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/amenities"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/member"
)

func RetrieveTableByCriteria(criteria map[string]string) ([]bukakoneksi.Table, error) {
	tables := make([]bukakoneksi.Table, 0)
	var err error

	query := "select * from tables"
	queryOrder := "order by name asc"

	// Query Filter
	queryFilter := ""
	queryArgs := []interface{}{}

	officeFloorID := criteria["office_floor_id"]
	name := criteria["name"]
	emptyWorkspace := criteria["empty_workspace"]
	memberTelegram := criteria["member_telegram"]
	source := criteria["source"]

	if officeFloorID != "" {
		queryFilter = queryFilter + " and office_floor_id = ?"
		queryArgs = append(queryArgs, officeFloorID)
	}

	if name != "" {
		queryFilter = queryFilter + " and lower(name) like concat('%',?,'%')"
		queryArgs = append(queryArgs, name)
	}

	if emptyWorkspace != "" {
		queryFilter = queryFilter + " and empty_workspace > 0"
	}

	if memberTelegram != "" {
		var findMember bukakoneksi.Member
		err = bukakoneksi.Dbmap.SelectOne(&findMember, "select * from members where telegram_username = ?", memberTelegram)
		if findMember.ID == 0 {
			return tables, nil
		}
		queryFilter = queryFilter + " and id in (select table_id from workspaces where member_id = ?)"
		queryArgs = append(queryArgs, findMember.ID)
	}

	if queryFilter != "" {
		queryFilter = strings.Replace(queryFilter, "and", "where", 1)
	}

	query = fmt.Sprintf("%s %s %s", query, queryFilter, queryOrder)

	_, err = bukakoneksi.Dbmap.Select(&tables, query, queryArgs...)

	for i := range tables {
		var workspaces []bukakoneksi.Workspace
		_, err = bukakoneksi.Dbmap.Select(&workspaces, "select * from workspaces where table_id = ? order by position asc", tables[i].ID)
		tables[i].Workspaces = workspaces

		for j := range workspaces {
			var m bukakoneksi.Member
			err = bukakoneksi.Dbmap.SelectOne(&m, "select * from members where id = ?", workspaces[j].MemberID)
			m.PictureURL = member.GetMemberImageUrl(m.PictureImage)
			workspaces[j].Member = m
			workspaces[j].Amenities = amenities.ConvertStringToArray(workspaces[j].AmenitiesTags)
		}

		if source == "bot" {
			var of bukakoneksi.OfficeFloor
			var o bukakoneksi.Office

			err = bukakoneksi.Dbmap.SelectOne(&of, "select * from office_floors where id = ?", tables[i].OfficeFloorID)
			err = bukakoneksi.Dbmap.SelectOne(&o, "select * from offices where id = ?", of.OfficeID)
			tables[i].OfficeFloor = &of
			tables[i].Office = &o
		} else {
			tables[i].OfficeFloor = nil
			tables[i].Office = nil
		}
	}

	return tables, err
}
