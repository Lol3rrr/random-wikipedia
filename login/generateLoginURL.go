package login

import "net/url"

func generateLoginURL(baseURL, email, password string) string {
	escapedEmail := url.QueryEscape(email)
	escapedPassword := url.QueryEscape(password)

	return baseURL + "/login/confirm?email=" + escapedEmail + "&password=" + escapedPassword
}
