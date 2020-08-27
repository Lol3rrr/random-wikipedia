package wikipedia

// list holds some info about the list and articles listed in it
type list struct {
	ID       int
	Title    string
	Articles []Article
}

// Article represents a single wikipedia article
type Article struct {
	ID    int    `json:"pageid"`
	Title string `json:"title"`
	URL   string `json:"fullurl"`
}

// List is a simple abstraction of a list to allow for easier testing
type List interface {
	// GetID returns the Page-ID of the List
	GetID() int
	// GetTitle returns the Title of the List
	GetTitle() string
	// GetArticles returns all the Articles in the List
	GetArticles() []Article
	// GetRandomArticle returns a single, random article from the List
	GetRandomArticle() (Article, error)
}

// Session is a simple abstraction to allow for easier testing
type Session interface {
	// GetArticlesInList simply returns an array of articles present in a given List
	GetArticlesInList(pageID int, plContinue string) ([]Article, string, error)
	// GetList loads all articles in the list
	GetList(pageID int, title string) (List, error)
}

// session is the internal actual struct for a sessionb
type session struct {
	BaseURL   string
	UserAgent string
}
