package util

import "net/url"

func IsValidURL(inputURL string) bool {
	parsedURL, err := url.ParseRequestURI(inputURL)
	if err != nil {
		return false
	}

	// Additional checks to ensure it's a full URL
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return false
	}

	return true
}
