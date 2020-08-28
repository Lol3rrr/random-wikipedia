package login

import (
	"net/url"
	"path"
)

func generateLoginURL(rawBaseURL, email, password string) string {
	escapedEmail := url.QueryEscape(email)
	escapedPassword := url.QueryEscape(password)

	baseURL, _ := url.Parse(rawBaseURL)
	baseURL.Path = path.Join(baseURL.Path, "/login/confirm")
	baseURL.Query().Add("email", escapedEmail)
	baseURL.Query().Add("password", escapedPassword)

	return baseURL.String()
}
