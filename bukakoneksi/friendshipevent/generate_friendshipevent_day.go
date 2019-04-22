package friendshipevent

import "time"

func GenerateFriendshipEventDay(day string) string {
	dateformat := "2006-01-02"
	if day != "" {
		_, err := time.Parse(dateformat, day)
		if err != nil {
			day = ""
		}
	}

	if day == "" {
		loc, _ := time.LoadLocation("Asia/Jakarta")
		day = time.Now().In(loc).Format(dateformat)
	}

	return day
}
