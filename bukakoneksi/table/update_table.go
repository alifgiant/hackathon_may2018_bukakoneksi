package table

import (
	"log"

	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/workspace"

	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi"
)

func TableExist(t bukakoneksi.Table) (bukakoneksi.Table, error) {
	var table bukakoneksi.Table

	err := bukakoneksi.Dbmap.SelectOne(&table, "select * from tables where id = ?", t.ID)
	if err != nil {
		log.Println(err)
		return table, err
	}

	return table, nil
}

func UpdateTable(t, tableEdit bukakoneksi.Table) (bukakoneksi.Table, error) {
	tableEdit.Name = t.Name
	tableEdit.WorkspaceSize = t.WorkspaceSize
	tableEdit.EmptyWorkspace = t.WorkspaceSize - len(t.Workspaces)

	// delete workspace
	_, err := bukakoneksi.Transaction.Exec("delete from workspaces where table_id = ?", t.ID)
	if err != nil {
		log.Println(err)
		return tableEdit, err
	}

	// recreate workspace
	var newWorkspaces []bukakoneksi.Workspace
	for _, w := range t.Workspaces {
		w.TableID = t.ID
		newW, err := workspace.CreateWorkspace(w)
		if err != nil {
			log.Println(err)
			return tableEdit, err
		}
		newWorkspaces = append(newWorkspaces, newW)
	}

	// update table
	tableEdit.Workspaces = newWorkspaces
	_, err = bukakoneksi.Transaction.Update(&tableEdit)
	if err != nil {
		log.Println(err)
		return tableEdit, err
	}

	return tableEdit, nil
}
