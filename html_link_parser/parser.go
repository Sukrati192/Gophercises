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
	var links []Links
	links = getTextAndLinkFromHtml(doc)
	return links
}

func isAnchorTag(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "a"
}

func getTextAndLinkFromHtml(n *html.Node) []Links {
	links := []Links{}
	if isAnchorTag(n) {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				textArr := []string{}
				for child := n.FirstChild; child != nil; child = child.NextSibling {
					if child == nil {
						break
					}

					t := addText(child)
					if t == nil {
						continue
					}
					textArr = append(textArr, *t)
				}
				text := strings.Join(textArr, "")
				link := Links{attr.Val, strings.TrimSpace(text)}
				links = append(links, link)
			}
		}
		return links
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, getTextAndLinkFromHtml(c)...)
	}
	return links
}

func addText(n *html.Node) *string {
	if n == nil {
		return nil
	}
	if n.Type == html.TextNode {
		str := n.Data
		var frontSpacesStrippedText string
		if frontSpacesStrippedText = str; strings.HasPrefix(str, "\n") {
			frontSpacesStrippedText = strings.TrimLeft(str, "\n ")
		}

		spacesStrippedText := strings.TrimRight(frontSpacesStrippedText, "\n ")

		if spacesStrippedText == frontSpacesStrippedText {
			// Nothing after text chars
			return &spacesStrippedText
		}

		retVal := spacesStrippedText
		for i := len(spacesStrippedText); i < len(frontSpacesStrippedText) && frontSpacesStrippedText[i] != byte(10); i++ {
			retVal += string(frontSpacesStrippedText[i])
		}

		return &retVal
	}
	return addText(n.FirstChild)
}

func readFile(inputFile string) io.Reader {
	fileName := fmt.Sprint("./examples/", inputFile)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Coudn't open input file. Error: ", err)
	}
	return file
}

/***
- This is how the text nodes look like for various test cases: 2, 3, 4, 7, 8
- I have printed string, size, json encoded string (to see all chars)
- The pattern is this, if string starts with \n strip all spaces from front
- Strip everything folowing the first \n after last char in text Eg. "chck \n    " -> "chck "


A link to another page 22 "A link to another page"

            Check me out on twitter
             49 "\n            Check me out on twitter\n            "

         9 "\n        "

            Gophercises is on  31 "\n            Gophercises is on "
Github 6 "Github"
!
         10 "!\n        "
Login  6 "Login "
Lost? Need help? 16 "Lost? Need help?"
@marcusolsson 13 "@marcusolsson"
dog cat
         16 "dog cat\n        "

     5 "\n    "
text inside dog link 20 "text inside dog link"

     5 "\n    "
Something in a span 19 "Something in a span"
 Text not in a span
     24 " Text not in a span\n    "
Bold text! 10 "Bold text!"

 1 "\n"


*/
