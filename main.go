package main

import (
	"fmt"
	"strings"

	"github.com/anaskhan96/soup"
)

func main() {
	//resp, err := soup.Get("https://en.wikipedia.org/w/index.php?search=QUERY&title=Special:Search&fulltext=Search&ns0=1")
	//if err != nil {
	//panic(err)
	//}

	//doc := soup.HTMLParse(resp).
	//Find("body").
	//Find("div", "id", "content").
	//Find("div", "id", "bodyContent").
	//Find("div", "id", "mw-content-text").
	//Find("div", "class", "searchresults").
	//Find("ul", "class", "mw-search-results").
	//FindAll("li", "class", "mw-search-result")
	////content := doc.Find("div", "id", "content")
	////bodyContent := content.Find("div", "id", "bodyContent")
	////mwContent := bodyContent.Find("div", "id", "mw-content-text")

	//for _, e := range doc {
	//heading := e.Find("div", "class", "mw-search-result-heading").Find("a")
	//link := heading.Attrs()["href"]
	//title := heading.Attrs()["title"]
	//metaData := e.Find("div", "class", "mw-search-result-data").Text()

	//previewText := e.Find("div", "class", "searchresult").FullText()

	//fmt.Println(title)
	//fmt.Println(link)
	//fmt.Println(metaData)
	//fmt.Println(previewText)
	//fmt.Println()

	//}

	//fmt.Println(len(doc))

	res := SearchArticles("query")
	for _, e := range res {
		fmt.Println(e.Title)
	}
}

func SearchArticles(searchPhrase string) []SearchResult {
	// replace space with plus to allow interpolation in URL
	formatted := strings.ReplaceAll(searchPhrase, " ", "+")
	// get the HTML
	resp, err := soup.Get(fmt.Sprintf("https://en.wikipedia.org/w/index.php?search=%s&title=Special:Search&fulltext=Search&ns0=1", formatted))
	if err != nil {
		panic(err)
	}

	// parse out the list of results
	searchResults := soup.HTMLParse(resp).
		Find("body").
		Find("div", "id", "content").
		Find("div", "id", "bodyContent").
		Find("div", "id", "mw-content-text").
		Find("div", "class", "searchresults").
		Find("ul", "class", "mw-search-results").
		FindAll("li", "class", "mw-search-result")

	results := make([]SearchResult, len(searchResults), len(searchResults))

	for i, e := range searchResults {
		res := SearchResult{}
		// parse out necessary info from each result
		heading := e.Find("div", "class", "mw-search-result-heading").Find("a")
		res.Link = heading.Attrs()["href"]
		res.Title = heading.Attrs()["title"]
		res.Metadata = e.Find("div", "class", "mw-search-result-data").Text()
		res.PreviewText = e.Find("div", "class", "searchresult").FullText()

		// add result to results lists
		results[i] = res
	}

	return results
}

type SearchResult struct {
	Title       string
	Link        string
	Metadata    string
	PreviewText string
}
