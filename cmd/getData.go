package cmd

import (
	"io"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

func GetSchema() UserstylesRoot {
	url := "https://raw.githubusercontent.com/catppuccin/userstyles/main/scripts/userstyles.yml"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	userstylesYml, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var data UserstylesRoot
	yaml.Unmarshal([]byte(userstylesYml), &data)

	return data
}
