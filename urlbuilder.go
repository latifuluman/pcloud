package pcloud

import "net/url"

// urlBuilder; return url with GET-params
func urlBuilder(method string, values url.Values, isEU bool) string {
	const (
		apiScheme = "https"
	)

	var apiHost string
	if isEU == true {
		apiHost = "eapi.pcloud.com"
	} else {
		apiHost = "api.pcloud.com"
	}

	u := url.URL{
		Scheme:   apiScheme,
		Host:     apiHost,
		Path:     method,
		RawQuery: values.Encode(),
	}
	return u.String()
}
