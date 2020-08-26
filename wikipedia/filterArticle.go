package wikipedia

import "strings"

func filterArticle(title string) bool {
	if strings.Index(title, "Template:") == 0 {
		return false
	}
	if strings.Index(title, "Template talk:") == 0 {
		return false
	}
	if strings.Index(title, "Portal:") == 0 {
		return false
	}
	if strings.Index(title, "Category:") == 0 {
		return false
	}
	if strings.Index(title, "Wikipedia:") == 0 {
		return false
	}
	if strings.Index(title, "Talk:") == 0 {
		return false
	}

	return true
}
