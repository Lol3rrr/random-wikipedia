package general

// User represents a single User instance
type User struct {
	ID           string   `json:"ID"`
	Subscription string   `json:"Subscription"`
	Settings     Settings `json:"Settings"`
	Lists        []int    `json:"Lists"`
}

// Settings reprents a single users Settings
type Settings struct {
	NotificationTime int `json:"NotificationTime"`
}

// ArticleNotification holds all the information for a Notification
type ArticleNotification struct {
	Title string `json:"Title"`
	URL   string `json:"URL"`
}
