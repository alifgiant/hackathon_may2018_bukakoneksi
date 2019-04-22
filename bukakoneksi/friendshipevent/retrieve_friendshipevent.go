package friendshipevent

import (
	"log"

	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi"
)

func RetrieveFriendshipEvent(city string, event string, day string) ([]bukakoneksi.FriendshipEvent, error) {
	var fe []bukakoneksi.FriendshipEvent

	day = GenerateFriendshipEventDay(day)

	_, err := bukakoneksi.Dbmap.Select(&fe, "select * from friendship_events where city = ? and event = ? and day = ?", city, event, day)
	if err != nil {
		log.Println(err)
		return fe, err
	}

	return fe, nil
}
