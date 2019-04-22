package member

import "github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/utils"

func GetMemberImageUrl(picture_image string) string {
	if picture_image == "" {
		picture_image = "default.jpg"
	}
	return utils.GetPic(picture_image)
}
