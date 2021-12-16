package main

import (
	"fmt"

	"github.com/Sukrati192/Gophercises/html_link_parser/parser"
)

func main() {
	var inputFile string
	fmt.Printf("Enter the html file to be parsed.\n")
	fmt.Scanln(&inputFile)
	links := parser.ParseLinkAndTextFromHtml(inputFile)
	fmt.Println(links)
}
