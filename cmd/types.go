package cmd

type UserstylesRoot struct {
	Collaborators []Collaborator       `yaml:"collaborators"`
	Userstyles    map[string]Userstyle `yaml:"userstyles"`
}

type Collaborator struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

type Userstyle struct {
	Name     interface{} `yaml:"name"`
	Category []string    `yaml:"category"`
	Icon     string      `yaml:"icon"`
	Color    string      `yaml:"color"`
	Readme   Readme      `yaml:"readme"`
}

type Readme struct {
	AppLink            []string       `yaml:"app-link"`
	Usage              string         `yaml:"usage"`
	Faq                []Faq          `yaml:"faq"`
	CurrentMaintainers []Collaborator `yaml:"current-maintainers"`
	PastMaintainers    []Collaborator `yaml:"past-maintainers"`
}

type Faq struct {
	Question string `yaml:"question"`
	Answer   string `yaml:"answer"`
}
