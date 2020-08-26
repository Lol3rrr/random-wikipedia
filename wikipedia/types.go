package wikipedia

// Article represents a single wikipedia article
type Article struct {
	ID    int    `json:"pageid"`
	Title string `json:"title"`
	URL   string `json:"fullurl"`
}

// Session is a simple abstraction to allow for easier testing
type Session interface {
	// GetArticlesInList simply returns an array of articles present in a given List
	GetArticlesInList(pageID int, plContinue string) ([]Article, string, error)
	// GetAllArticlesInList loads all articles in the list
	GetAllArticlesInList(pageID int) ([]Article, error)
}

// session is the internal actual struct for a sessionb
type session struct {
	BaseURL   string
	UserAgent string
}
