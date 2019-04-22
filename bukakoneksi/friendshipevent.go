package bukakoneksi

// FriendshipEvent struct
type FriendshipEvent struct {
	ID               int    `db:"id" json:"id"`
	Event            string `db:"event" json:"event"` // lunch or dinner
	TelegramUserID   string `db:"telegram_user_id" json:"telegram_user_id"`
	TelegramUsername string `db:"telegram_username" json:"telegram_username"`
	City             string `db:"city" json:"city"`
	Day              string `db:"day" json:"day"` // String yyyy-mm-dd
}
