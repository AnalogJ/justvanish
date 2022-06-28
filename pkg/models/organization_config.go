package models

type OrganizationConfig struct {
	OrganizationName string `yaml:"organization_name"`
	Website          string `yaml:"website"`
	Contact          struct {
		Mail []struct {
			Address string   `yaml:"address"`
			Usage   []string `yaml:"usage"`
		} `yaml:"mail"`
		Email []struct {
			Address string   `yaml:"address"`
			Usage   []string `yaml:"usage"`
		} `yaml:"email"`
	} `yaml:"contact"`
	OrganizationType []string `yaml:"organization_type"`
	Regulation       []string `yaml:"regulation"`
	Notes            string   `yaml:"notes"`
}
