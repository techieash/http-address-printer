package util

import "net/url"

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func ConvertToURL(str string) string {
	u, _ := url.Parse(str)
	u.Scheme = "http"
	return u.String()
}
