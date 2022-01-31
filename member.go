package ghanamps

import (
	"fmt"
	"path"
	"strings"

	"github.com/gocolly/colly"
)

// Member represents the details of one member of parliament
type Member struct {
	Name         string `json:"name"`
	Party        string `json:"party"`
	Constituency string `json:"constituency"`
	Region       string `json:"region"`
	Photo        string `json:"photo"`
	Profile      string `json:"profile"`
}

// baseURL is the base url to the website of Ghana's parliament
const baseURL string = "https://www.parliament.gh"

// Fetch fetches the current members of parliament
func Fetch() ([]Member, error) {
	members := []Member{}
	c := colly.NewCollector()
	pageVisitor := colly.NewCollector(colly.Async(true))

	c.OnHTML("a.square", func(e *colly.HTMLElement) {
		pageVisitor.Visit(baseURL + "/" + e.Attr("href"))
	})

	pageVisitor.OnHTML("div.mpcard", func(e *colly.HTMLElement) {
		member := Member{}
		member.Name = strings.TrimSpace(strings.ToUpper(e.ChildText("b.padd")))
		member.Constituency = strings.TrimSpace(strings.ToUpper(e.ChildText("a > div > center b:nth-of-type(2)")))
		member.Region = strings.TrimSpace(strings.ToUpper(e.ChildText("a > div > center span:nth-of-type(1)")))
		member.Photo = strings.TrimSpace(baseURL + "/" + e.ChildAttr("a > img", "src"))
		member.Party = strings.TrimSpace(e.ChildText("a > div > center :nth-child(2)"))

		mpNumber := strings.TrimSpace(strings.TrimSuffix(path.Base(member.Photo), path.Ext(member.Photo)))
		member.Profile = baseURL + fmt.Sprintf("/mps?mp=%s", mpNumber)

		members = append(members, member)
	})

	c.Visit(baseURL + "/mps?az")
	pageVisitor.Wait()
	return members, nil
}
