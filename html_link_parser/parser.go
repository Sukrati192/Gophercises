package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Links struct {
	Href string
	Text string
}

func main() {
	var inputFile string
	fmt.Printf("Enter the html file to be parsed.\n")
	fmt.Scanln(&inputFile)
	links := ParseLinkAndTextFromHtml(inputFile)
	fmt.Print(links)
}

func ParseLinkAndTextFromHtml(inputfile string) []Links {
	file := readFile(inputfile)
	doc, err := html.Parse(file)
	if err != nil {
		log.Fatal(err)
	}
	links := []Links{}
	links = getTextAndLinkFromHtml(doc, links)
	return links
}

func getTextAndLinkFromHtml(n *html.Node, links []Links) []Links {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				text := ""
				for child := n.FirstChild; child != nil; child = child.NextSibling {
					text = addText(child, text)
				}
				link := Links{a.Val, strings.TrimSpace(text)}
				links = append(links, link)
			}
		}
		return links
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = getTextAndLinkFromHtml(c, links)
	}
	return links
}

func addText(n *html.Node, text string) string {
	if n == nil {
		if strings.TrimSpace(text) == text {
			return text
		}
		return (strings.TrimSpace(text) + " ")
	}
	if n.Type == html.TextNode {
		data := n.Data
		if strings.TrimSpace(text) != text {
			text = strings.TrimSpace(text) + " "
		}
		if strings.TrimSpace(n.Data) != n.Data {
			data = strings.TrimSpace(n.Data) + " "
		}
		return text + data
	}
	return addText(n.FirstChild, text)
}

func readFile(inputFile string) io.Reader {
	fileName := fmt.Sprint("/home/sukrati/Gophercises/html_link_parser/examples/", inputFile)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Coudn't open input file. Error: ", err)
	}
	return file
}
