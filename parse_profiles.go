package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/vaughan0/go-ini"
)

func parseProfiles() ([]string, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	configFile, err := ini.LoadFile(fmt.Sprintf("%s/.aws/config", dirname))

	if err != nil {
		fmt.Println("❌ No or malformatted AWS configuration found")
		return nil, err
	}

	var profileNames = []string{}
	for k := range configFile {
		profileNames = append(profileNames, strings.Replace(k, "profile ", "", 1))
	}

	sort.Strings(profileNames)

	return profileNames, nil
}
