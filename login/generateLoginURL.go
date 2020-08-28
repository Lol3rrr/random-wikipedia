package login

import (
	"net/url"
	"path"
)

func generateLoginURL(rawBaseURL, email, password string) string {
	baseURL, _ := url.Parse(rawBaseURL)
	baseURL.Path = path.Join(baseURL.Path, "/login/confirm")

	q, _ := url.ParseQuery("")
	q.Add("email", email)
	q.Add("password", password)
	baseURL.RawQuery = q.Encode()

	return baseURL.String()
}
