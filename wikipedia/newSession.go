package wikipedia

// NewSession is used to obtain a new Session
func NewSession() Session {
	return &session{
		BaseURL:   "https://en.wikipedia.org/",
		UserAgent: "Random-Wikipedia/0.1 (;leon@lol3r.net)",
	}
}
