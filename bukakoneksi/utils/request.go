package utils

import "net/url"

func IsQueryParamExist(params url.Values, name string) bool {
	if params[name] == nil || len(params[name][0]) == 0 {
		return false
	}
	return true
}
