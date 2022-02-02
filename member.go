package ghanamps

import (
	"fmt"
	"path"
	"strings"

	"github.com/gocolly/colly"
)

// Member represents the details of one member of parliament
type Member struct {
	ID           string `json:"-"`
	Name         string `json:"name"`
	Party        string `json:"party"`
	PartyAbbr    string `json:"partyAbbr"`
	Constituency string `json:"constituency"`
	Region       string `json:"region"`
	Photo        string `json:"photo"`
	Profile      string `json:"profile"`
}

// baseURL is the base url to the website of Ghana's parliament
const baseURL string = "https://www.parliament.gh"

// Fetch fetches the current members of parliament
func Fetch() ([]Member, error) {
	members := map[string]Member{}
	payload := []Member{}

	mainCollector := colly.NewCollector()
	pageVisitor := colly.NewCollector(colly.Async(true))
	profileVisitor := colly.NewCollector(colly.Async(true))

	mainCollector.OnHTML("a.square", func(e *colly.HTMLElement) {
		pageVisitor.Visit(baseURL + "/" + e.Attr("href"))
	})

	pageVisitor.OnHTML("div.mpcard", func(e *colly.HTMLElement) {
		member := Member{}
		member.Photo = strings.TrimSpace(baseURL + "/" + e.ChildAttr("a > img", "src"))
		member.Name = strings.TrimSpace(strings.ToUpper(e.ChildText("b.padd")))
		member.Constituency = strings.TrimSpace(strings.ToUpper(e.ChildText("a > div > center b:nth-of-type(2)")))
		member.Region = strings.TrimSpace(strings.ToUpper(e.ChildText("a > div > center span:nth-of-type(1)")))
		member.PartyAbbr = strings.TrimSpace(e.ChildText("a > div > center :nth-child(2)"))

		memberNumber := strings.TrimSpace(strings.TrimSuffix(path.Base(member.Photo), path.Ext(member.Photo)))
		member.Profile = baseURL + fmt.Sprintf("/mps?mp=%s", memberNumber)

		members[memberNumber] = member
		profileVisitor.Visit(member.Profile)
	})

	profileVisitor.OnHTML("div.fl > table", func(e *colly.HTMLElement) {
		partyParts := strings.Split(e.ChildText("tr:nth-child(2) > td:nth-child(2)"), "(")
		memberNumber := e.Request.URL.Query().Get("mp")
		member := members[memberNumber]
		if len(partyParts) > 1 {
			member.Party = strings.TrimSpace(strings.ToUpper(partyParts[0]))
			payload = append(payload, member)
		}
	})

	mainCollector.Visit(baseURL + "/mps?az")
	pageVisitor.Wait()
	profileVisitor.Wait()
	return payload, nil
}

// Members is an alias of Fetch.
// It returns the current members of parliament
func Members() ([]Member, error) {
	return Fetch()
}
