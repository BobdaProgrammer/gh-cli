package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"

	"github.com/nsf/termbox-go"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type Size struct {
	Width  int
	Height int
}

var size Size
var AccessToken string

type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stars       int    `json:"stargazers_count"`
	Language    string `json:"language"`
}

func init() {
	err := termbox.Init()
	if err != nil {
		panic("Your terminal is not working for us currently.")
	}
	var width, height int = termbox.Size()
	size.Width = width
	size.Height = height
}
func repos(username string) {
	url := "https://api.github.com/users/" + username + "/repos"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic("couldn't fetch repositories")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic("couldn't fetch repositories")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("couldn't read response body")
	}

	var repositories []Repository
	err = json.Unmarshal(body, &repositories)
	if err != nil {
		panic("couldn't parse JSON response")
	}
	nameWidth := int(math.Round(float64(size.Width / 6.0)))
	descWidth := int(math.Round(float64(size.Width / 2.0)))
	starWidth := int(math.Round(float64(size.Width / 22.0)))
	langWidth := int(math.Round(float64(size.Width / 12.0)))
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().
			Foreground(lipgloss.Color("99"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			return lipgloss.NewStyle().PaddingRight(1)
		}).
		Headers("NAME", "DESCRIPTION", "STARS", "LANGUAGE")
	for _, repo := range repositories {
		var name string = repo.Name
		if len(name) > nameWidth {
			name = name[:nameWidth-1] + "…"
		}
		var desc string = repo.Description
		if len(desc) > descWidth {
			desc = desc[:descWidth-1] + "…"
		}
		var star string = fmt.Sprint(repo.Stars)
		if len(star) > starWidth {
			star = star[:starWidth-1] + "…"
		}
		var lang string = repo.Language
		if len(lang) > langWidth {
			lang = lang[:langWidth-1] + "…"
		}
		t.Row(name, desc, star, lang)
	}
	fmt.Println(t)
}

func followers(username string) []string {
	return nil
}
func following(username string) []string {
	return nil
}
func issues(repo string) []string {
	return nil
}
func main() {
	options := os.Args[1:]
	flag := options[0]
	switch flag {
	case "-r": //repos
		name := options[1]
		fmt.Println(name + "'s repos:")
		repos(name)
	case "-fr":
		name := options[1]
		fmt.Println(name + "'s followers")
	case "-fi":
		name := options[1]
		fmt.Println(name + "'s following")
	case "-i":
		repo := options[1]
		fmt.Println(repo + "'s issues")
	}
}
