package pkg

type Experience struct {
	Role     string   `yaml:"Role"`
	Company  string   `yaml:"Company"`
	Start    string   `yaml:"Start"`
	End      string   `yaml:"End"`
	Location string   `yaml:"Location"`
	Clients  []string `yaml:"Clients"`
	Details  []string `yaml:"Details"`
}

type Resume struct {
	Name        string      `yaml:"Name"`
	Role        string      `yaml:"MyRole"`
	Email       string      `yaml:"Email"`
	Phone       string      `yaml:"Phone"`
	Location    string      `yaml:"Location"`
	PageTitle   string      `yaml:"PageTitle"`
	Summary     string      `yaml:"Summary"`
	Education   []Education `yaml:"Education"`
	LinkedInUrl string      `yaml:"LinkedInUrl"`

	Experience []Experience `yaml:"Experience"`
	Projects   []string     `yaml:"Projects"`
	Languages  []string     `yaml:"Languages"`
	Skills     []Skills     `yaml:"Skills"`
	Keywords   []string     `yaml:"Keywords"`
}

type Education struct {
	School string `yaml:"School"`
	Degree string `yaml:"Degree"`
	Major  string `yaml:"Major"`
	Year   string `yaml:"Year"`
}

type Skills struct {
	Name  string   `yaml:"Name"`
	Items []string `yaml:"Items"`
}
