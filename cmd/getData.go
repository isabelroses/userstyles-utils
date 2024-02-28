package cmd

import (
	"io"
	"log"
	"net/http"
	"strings"

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

func InferCollaboratorName(collaborator Collaborator) string {
	var inferedName string

	if collaborator.Name == "" {
		inferedName = strings.Replace(collaborator.Url, "https://github.com/", "", -1)
	} else {
		inferedName = collaborator.Name
	}

	return inferedName
}
