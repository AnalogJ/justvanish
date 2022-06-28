package models

type UserConfig struct {
	FirstName      string   `yaml:"first_name" mapstructure:"first_name"`
	LastName       string   `yaml:"last_name" mapstructure:"last_name"`
	EmailAddresses []string `yaml:"email_addresses" mapstructure:"email_addresses"`
	MailAddresses  []string `yaml:"mail_addresses" mapstructure:"mail_addresses"`
	PhoneNumbers   []string `yaml:"phone_numbers" mapstructure:"phone_numbers"`
}
