package pkg

type Resume struct {
	Name      string `yaml:"Name"`
	Role      string `yaml:"MyRole"`
	Email     string `yaml:"Email"`
	Location  string `yaml:"Location"`
	PageTitle string `yaml:"PageTitle"`
	Mission   string `yaml:"Mission"`

	Experience []Experience `yaml:"Experience"`
	Projects   []string     `yaml:"Projects"`
	Languages  []string     `yaml:"Languages"`
	Skills     []string     `yaml:"Skills"`
	Keywords   []string     `yaml:"Keywords"`
}

type Experience struct {
	Role     string   `yaml:"Role"`
	Company  string   `yaml:"Company"`
	Start    string   `yaml:"Start"`
	End      string   `yaml:"End"`
	Location string   `yaml:"Location"`
	Details  []string `yaml:"Details"`
}
