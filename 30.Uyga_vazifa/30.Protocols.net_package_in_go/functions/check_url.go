package functions

import "net/url"

func IsValidURL(input string) bool {
	_, err := url.ParseRequestURI(input)
	return err == nil
}
