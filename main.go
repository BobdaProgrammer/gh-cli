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
func widths(proportions []float64,text []string)[]string{
	var out []string;
	for i,proportion := range proportions{
		width:=int(math.Round(float64(size.Width) / proportion))
		val := text[i]
		if len(val)>width{
			val = val[:width-1]+"…"
		}
		out=append(out, val)
	}
	return out
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
	nameWidth := 6.0
	descWidth := 1.75
	starWidth := 22.0
	langWidth := 12.0
	var Tablewidths []float64 = []float64{nameWidth, descWidth, starWidth, langWidth}
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().
			Foreground(lipgloss.Color("99"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			return lipgloss.NewStyle().PaddingRight(1)
		}).
		Headers("NAME", "DESCRIPTION", "STARS", "LANGUAGE")
	for _, repo := range repositories {
		var data []string = []string{repo.Name, repo.Description, fmt.Sprint(repo.Stars), repo.Language}
		var out []string = widths(Tablewidths, data)
		t.Row(out[0], out[1], out[2], out[3])
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
	case "-h":
		var help string = `
		commands:
			-h						help/instructions
			-r <username>			list of repos from username
			-fr <username>			list of followers from username
			-fi <username>			list of following from username
			-i <username>/<repo>	list of issues from repo
		`
		fmt.Print(help)
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
