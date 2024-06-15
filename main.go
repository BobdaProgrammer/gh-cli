package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().
			Foreground(lipgloss.Color("99"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			return lipgloss.NewStyle().PaddingRight(1)
		}).
		Headers("NAME", "DESCRIPTION", "STARS", "LANGUAGE")
	t.Row("iew0jfe0rjfer0jgeroigmnroeignmreiomgoerimgoerimgoremgrieg", "oiwefmneifmeorifmeroigmerigmeimgoerimgreoimgeoimgeroigmreoigmeroigmeroimgoerigmreoimger", "9875640786098760", "suminanimiamaiai")
	for _, repo := range repositories {
		t.Row(repo.Name, repo.Description, fmt.Sprint(repo.Stars), repo.Language)
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
