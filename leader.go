package ghanamps

import (
	"strings"

	"github.com/gocolly/colly"
)

// Leader represents a leader of parliament
type Leader struct {
	ID    string `json:"-"`
	Name  string `json:"name"`
	Title string `json:"title"`
	Photo string `json:"photo"`
}

// Leaders returns a JSON data of the leadership of the parliament.
func Leaders() ([]Leader, error) {
	c := colly.NewCollector()
	leaders := []Leader{}

	c.OnHTML("div#content > table.gone > tbody > tr > td > center", func(e *colly.HTMLElement) {
		leader := Leader{}
		if len(e.DOM.Contents().Nodes) == 3 {
			leader.Name = strings.TrimSpace(e.DOM.Contents().Nodes[2].FirstChild.Data)
			leader.Title = strings.TrimSpace(e.DOM.Contents().Nodes[2].FirstChild.NextSibling.NextSibling.Data)
		}

		if len(e.DOM.Contents().Nodes) > 3 {
			leader.Name = strings.TrimSpace(e.DOM.Contents().Nodes[2].Data)
			leader.Title = strings.TrimSpace(e.DOM.Contents().Nodes[4].Data)
		}
		leader.Photo = baseURL + "/" + e.ChildAttr("img", "src")
		leaders = append(leaders, leader)
	})

	c.Visit(baseURL + "/mps?leadership")
	return leaders, nil
}
