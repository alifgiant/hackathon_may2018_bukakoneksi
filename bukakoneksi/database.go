package bukakoneksi

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

var Dbmap *gorp.DbMap
var Db *sql.DB

var Transaction *gorp.Transaction

func InitDB(dsn string) error {
	Db, _ = sql.Open("mysql", fmt.Sprintf("%s?charset=utf8&parseTime=True&loc=Local", dsn))
	err := Db.Ping()
	if err != nil {
		log.Println(err)
		return err
	}

	database := &gorp.DbMap{Db: Db, Dialect: gorp.MySQLDialect{}}

	database.AddTableWithName(Amenities{}, "amenities").SetKeys(true, "ID")
	database.AddTableWithName(City{}, "cities").SetKeys(true, "ID")
	database.AddTableWithName(Member{}, "members").SetKeys(true, "ID")
	database.AddTableWithName(Office{}, "offices").SetKeys(true, "ID")
	database.AddTableWithName(OfficeFloor{}, "office_floors").SetKeys(true, "ID")
	database.AddTableWithName(Table{}, "tables").SetKeys(true, "ID")
	database.AddTableWithName(Workspace{}, "workspaces").SetKeys(true, "ID")
	database.AddTableWithName(FriendshipEvent{}, "friendship_events").SetKeys(true, "ID")
	Dbmap = database
	return nil
}
