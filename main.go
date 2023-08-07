package main

import (
	"log"
	"os"

	"github.com/lithammer/fuzzysearch/fuzzy"

	"github.com/manifoldco/promptui"
)

var version string = "v1.1.0"

func main() {
	profiles, err := parseProfiles()

	if err != nil {
		log.Fatal((err))
		return
	}

	/**
	 * In case the user already provided a profile
	 * to search for in the command line arguments,
	 * try to set it without using the prompt.
	 */
	if len(os.Args) > 1 {
		argument := os.Args[1]

		matchingProfiles := fuzzy.RankFind(argument, profiles)

		if matchingProfiles.Len() > 0 {
			setProfile(matchingProfiles[0].Target)
			return
		}
	}

	/**
	 * In case no profile could be found,
	 * or no argument has been provided,
	 * open the prompt.
	 */
	searcher := func(input string, index int) bool {
		item := profiles[index]
		return fuzzy.MatchFold(input, item)
	}

	templates := &promptui.SelectTemplates{
		Active:   `{{ "> " | cyan | bold }}{{ . | cyan | bold }}`,
		Inactive: `  {{ . }}`,
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
