package general

// User represents a single User instance
type User struct {
	ID           string    `json:"ID"`
	Subscription string    `json:"Subscription"`
	Settings     Settings  `json:"Settings"`
	Lists        []List    `json:"Lists"`
	Favorites    []Article `json:"Favorites"`
}

// Settings reprents a single users Settings
type Settings struct {
	NotificationTime int `json:"NotificationTime"`
}

// Article holds all the information needed about a single article
type Article struct {
	ID    int    `json:"ID"`
	Title string `json:"Title"`
	URL   string `json:"URL"`
}

// List holds all the information about a single list
type List struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
}

// ArticleNotification holds all the information for a Notification
type ArticleNotification struct {
	ID    int    `json:"ID"`
	Title string `json:"Title"`
	URL   string `json:"URL"`
}
