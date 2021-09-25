package main

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

const baseURL string = "https://www.parliament.gh"

type Member struct {
	Name         string `json:"name"`
	PartyName    string `json:"partyName"`
	PartyAbbr    string `json:"partyAbbr"`
	Constituency string `json:"constituency"`
	Region       string `json:"region"`
	Photo        string `json:"photo"`
}

func main() {
	members := make([]Member, 275)
	c := colly.NewCollector()
	pageVisitor := colly.NewCollector(colly.Async(true))

	c.OnHTML("a.square", func(e *colly.HTMLElement) {
		pageVisitor.Visit(baseURL + "/" + e.Attr("href"))
	})

	index := 0
	pageVisitor.OnHTML("div.mpcard", func(e *colly.HTMLElement) {
		partyDetails := strings.Split(e.ChildText("a > div > center :nth-child(2)"), " (")
		members[index] = Member{
			Name:         strings.ToUpper(e.ChildText("b.padd")),
			PartyName:    strings.ToUpper(partyDetails[0]),
			PartyAbbr:    strings.ToUpper(partyDetails[1][:len(partyDetails[1])-1]),
			Constituency: strings.ToUpper(e.ChildText("a > div > center b:nth-of-type(2)")),
			Region:       strings.ToUpper(e.ChildText("a > div > center span:nth-of-type(1)")),
			Photo:        baseURL + "/" + e.ChildAttr("a > img", "src"),
		}
		index++
	})

	c.Visit(baseURL + "/mps?az")
	pageVisitor.Wait()

	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	enc.Encode(members)
}
