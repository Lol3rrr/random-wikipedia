package general

// User represents a single User instance
type User struct {
	ID           string
	Subscription string
	Settings     Settings
	Lists        []int
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
