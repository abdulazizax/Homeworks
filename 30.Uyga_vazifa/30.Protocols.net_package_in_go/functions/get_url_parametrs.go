package functions

import "net/url"

type URL struct {
	Schema       string
	Host         string
	Path         string
	Query_Params url.Values
	Fragment     string
}

func GetURL_Parametrs(raw_url string) (result URL, err error) {
	parsedURL, err := url.Parse(raw_url)
	if err != nil {
		return result, err
	}

	result.Schema = parsedURL.Scheme
	result.Host = parsedURL.Host
	result.Path = parsedURL.Path
	result.Query_Params = parsedURL.Query()
	result.Fragment = parsedURL.Fragment

	return result, err
}
