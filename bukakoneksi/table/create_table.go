package table

import (
	"log"

	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/amenities"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/member"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/workspace"
)

func CreateTable(t bukakoneksi.Table) (bukakoneksi.Table, error) {
	t.EmptyWorkspace = t.WorkspaceSize - len(t.Workspaces)
	t.Status = "unmap"

	var newWorkspaces []bukakoneksi.Workspace
	err := bukakoneksi.Transaction.Insert(&t)
	if err != nil {
		log.Println(err)
		return t, err
	}

	for _, w := range t.Workspaces {
		w.AmenitiesTags = amenities.ConvertArrayToString(w.Amenities)
		w.MemberID = w.MemberID

		m, err := member.RetrieveMemberByName(w.MemberName)
		if err != nil {
			log.Println(err)
			return t, err
		}
		w.Member = m
		w.TableID = t.ID

		newW, err := workspace.CreateWorkspace(w)
		if err != nil {
			log.Println(err)
			return t, err
		}
		newWorkspaces = append(newWorkspaces, newW)
	}

	t.Workspaces = newWorkspaces

	return t, nil
}
