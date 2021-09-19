package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

type options struct {
	Text string `json:"text,omitempty"`
	Arc  string `json:"arc,omitempty"`
}

type Pages struct {
	Title   string    `json:"title,omitempty"`
	Story   []string  `json:"story,omitempty"`
	Options []options `json:"options,omitempty"`
}

type story struct{}

var (
	storyPages   = parseJsonFile()
	htmlTemplate = template.Must(template.ParseFiles("layout.html"))
)

func parseJsonFile() map[string]Pages {
	jsonFile, err := os.Open("gopher.json")
	if err != nil {
		log.Fatalf("Could not open json file: ", err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	var pages map[string]Pages
	json.Unmarshal(byteValue, &pages)
	return pages
}

func (s story) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	htmlTemplate.Execute(w, storyPages[path])
}

func PrintStoryInTerminal(page string) {
	fmt.Println(storyPages[page].Title)
	for story := range storyPages[page].Story {
		fmt.Println(storyPages[page].Story[story])
	}
	if page == "home" {
		fmt.Println("Story is complete!!")
		return
	}
	for i := 1; i <= len(storyPages[page].Options); i++ {
		fmt.Println("Press", i, ": ", storyPages[page].Options[i-1].Text)
	}
	fmt.Println("Press 0 to Quit.")
	var option int
	fmt.Print("Choose an option: ")
	fmt.Scanln(&option)
	if option == 0 {
		return
	}
	for option > len(storyPages[page].Options) || option < 1 {
		fmt.Print("Invalid Option!!! Choose an option again: ")
		fmt.Scanln(&option)
		if option == 0 {
			return
		}
	}
	PrintStoryInTerminal(storyPages[page].Options[option-1].Arc)
}

func main() {
	fmt.Print("Mention the starting point for your story: ")
	var start string
	fmt.Scanln(&start)
	if _, ok := storyPages[start]; !ok {
		log.Fatalf("Invalid starting point!!")
	}
	PrintStoryInTerminal(start)
	http.Handle("/", story{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
