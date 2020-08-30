package wikipedia

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type queryResult struct {
	Pages map[string]Article `json:"pages"`
}

type linksResponse struct {
	BatchComplete string            `json:"batchcomplete"`
	Continue      map[string]string `json:"continue"`
	Limits        map[string]int    `json:"limits"`
	Query         queryResult       `json:"query"`
}

func (s *session) getArticlesInList(pageID int, plContinue string) ([]Article, string, error) {
	var urlBuilder strings.Builder
	urlBuilder.WriteString(s.BaseURL)
	urlBuilder.WriteString("/w/api.php?action=query&format=json&pageids=")
	urlBuilder.WriteString(strconv.Itoa(pageID))
	urlBuilder.WriteString("&generator=links&gpllimit=max&prop=info&inprop=url")
	if len(plContinue) > 0 {
		escapedContinue := url.QueryEscape(plContinue)

		urlBuilder.WriteString("&gplcontinue=")
		urlBuilder.WriteString(escapedContinue)
	}

	req, err := http.NewRequest("GET", urlBuilder.String(), nil)
	if err != nil {
		return nil, "", err
	}

	req.Header.Add("User-Agent", s.UserAgent)

	rawResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer rawResp.Body.Close()

	if rawResp.StatusCode != 200 {
		return nil, "", fmt.Errorf("Returned non 200 Status-Code: %v", rawResp.StatusCode)
	}

	body, err := ioutil.ReadAll(rawResp.Body)
	if err != nil {
		return nil, "", err
	}

	var resp linksResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, "", err
	}

	pages := make([]Article, 0, len(resp.Query.Pages))
	for _, value := range resp.Query.Pages {
		if !filterArticle(value.Title) {
			continue
		}

		pages = append(pages, value)
	}

	continueKey, found := resp.Continue["continue"]
	if found {
		continueKey = continueKey[:len(continueKey)-2]
	}

	return pages, resp.Continue["gplcontinue"], nil
}
