package bukakoneksi

// Member struct
type Member struct {
	ID               int    `db:"id" json:"id"`
	Name             string `db:"name" json:"name"`
	TelegramUserID   string `db:"telegram_user_id" json:"telegram_user_id"`
	TelegramUsername string `db:"telegram_username" json:"telegram_username"`
	PictureImage     string `db:"picture_image" json:"-"`
	PictureURL       string `db:"-" json:"picture_url"`
}
