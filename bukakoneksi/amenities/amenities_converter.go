package amenities

import (
	"strings"
)

func ConvertArrayToString(amenities_tags []string) string {
	res := ""
	for i := range amenities_tags {
		if i == 0 {
			res = res + amenities_tags[i]
		} else {
			res = res + "|" + amenities_tags[i]
		}
	}
	return res
}

func ConvertStringToArray(amenities_tags string) []string {
	if amenities_tags != "" {
		return strings.Split(amenities_tags, "|")
	} else {
		return make([]string, 0)
	}
}
