package friendshipevent

import (
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi"
)

func CreateFriendshipEvent(fe bukakoneksi.FriendshipEvent) (bukakoneksi.FriendshipEvent, error) {
	fe.Day = GenerateFriendshipEventDay(fe.Day)

	exist, err := bukakoneksi.Dbmap.SelectInt("select count(id) from friendship_events where event = ? and day = ? and telegram_user_id = ?", fe.Event, fe.Day, fe.TelegramUserID)
	if err != nil {
		return fe, err
	} else if exist == 0 {
		if err := bukakoneksi.Dbmap.Insert(&fe); err != nil {
			return fe, err
		}
	}

	return fe, nil
}
