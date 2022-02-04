package main

import "fmt"

var helpText = `
ghanamps - A cli to get data on the members and leadership of Ghana's parliament

USAGE
    ghanamps [SUBCOMMAND] [FLAG]

AVAILABLE SUBCOMMANDS
    members     get all members of parliament
    leaders     get all leaders of parliament

SUBCOMMAND FLAGS
    members
        --party=PARTY   filter members by party

EXAMPLES
    Eg 1:  Get all members
        ghanamps members
        
    Eg 2:  Get all members from ndc party
        ghanamps members --party ndc

    Eg 3:  Get all members from npp party
        ghanamps members --party npp
    
    Eg 4:  Get all independent members
        ghanamps members --party independent

    Eg 5:  Get all leaders
        ghanamps leaders

`

func usage() {
	fmt.Println(helpText)
}
