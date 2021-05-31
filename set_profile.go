package main

import (
	"fmt"
	"log"
	"os"
)

func setProfile(profile string) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(fmt.Sprintf("%s/.aws-profile-picker", dirname), []byte(profile), 0644)
}
