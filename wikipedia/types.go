package wikipedia

// Article represents a single wikipedia article
type Article struct {
	ID    int    `json:"pageid"`
	Title string `json:"title"`
	URL   string `json:"fullurl"`
}

// ListInfo holds basic information about a list
type ListInfo struct {
	ID    int    `json:"ID"`
	Title string `json:"Title"`
}

// list holds some info about the list and articles listed in it
type list struct {
	ListInfo
	Articles []Article
}

// Session is a simple abstraction to allow for easier testing
type Session interface {
	// UpdateLists is used to load the latest Version of all the registed lists
	UpdateLists()
	// GetRandomArticle is used to obtain a random Article out of all the
	// specified lists
	GetRandomArticle(listIDs []int) (Article, error)
	// GetLists simply returns a small list of all lists
	GetLists() []ListInfo
}

// session is the internal actual struct for a sessionb
type session struct {
	BaseURL   string
	UserAgent string
	Lists     map[int]list
}
