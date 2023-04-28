package main

import (
	"log"

	"github.com/lithammer/fuzzysearch/fuzzy"

	"github.com/manifoldco/promptui"
)

var version string = "v1.0.9"

func main() {
	profiles, err := parseProfiles()

	if err != nil {
		log.Fatal((err))
		return
	}

	templates := &promptui.SelectTemplates{
		Active:   `{{ "> " | cyan | bold }}{{ . | cyan | bold }}`,
		Inactive: `  {{ . }}`,
	}

	searcher := func(input string, index int) bool {
		item := profiles[index]
		return fuzzy.MatchFold(input, item)
	}

	prompt := promptui.Select{
		Label:             "Select AWS Profile",
		Items:             profiles,
		Size:              10,
		StartInSearchMode: true,
		Templates:         templates,
		Searcher:          searcher,
	}

	_, result, err := prompt.Run()

	if err != nil {
		log.Fatal(err)
		return
	}

	setProfile(result)
}
