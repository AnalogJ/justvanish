package models

import (
	"golang.org/x/exp/slices"
	"strings"
)

type OrganizationConfig struct {
	OrganizationName string                    `yaml:"organization_name"`
	Website          string                    `yaml:"website"`
	Contact          OrganizationConfigContact `yaml:"contact"`
	OrganizationType []string                  `yaml:"organization_type"`
	Regulation       []string                  `yaml:"regulation"`
	Notes            string                    `yaml:"notes"`
}

type OrganizationConfigContact struct {
	Mail  []OrganizationConfigContactInfo `yaml:"mail"`
	Email []OrganizationConfigContactInfo `yaml:"email"`
	Form  []OrganizationConfigContactInfo `yaml:"form"`
}

func (c *OrganizationConfigContact) AddMail(address string, usage []string) {
	address = strings.TrimSpace(address)
	if len(address) == 0 {
		return
	}

	if c.Mail == nil {
		c.Mail = []OrganizationConfigContactInfo{}
	}
	existing := false
	for _, contactInfo := range c.Mail {
		if contactInfo.Address == address {
			existing = true
			for _, usageStr := range usage {
				if !slices.Contains(contactInfo.Usage, usageStr) {
					contactInfo.Usage = append(contactInfo.Usage, usageStr)
				}
			}
		}
	}
	if !existing {
		c.Mail = append(c.Mail, OrganizationConfigContactInfo{Address: address, Usage: usage})
	}
}

func (c *OrganizationConfigContact) AddEmail(address string, usage []string) {
	address = strings.TrimSpace(address)
	if len(address) == 0 {
		return
	}

	if c.Email == nil {
		c.Email = []OrganizationConfigContactInfo{}
	}
	existing := false
	for _, contactInfo := range c.Email {
		if contactInfo.Address == address {
			existing = true
			for _, usageStr := range usage {
				if !slices.Contains(contactInfo.Usage, usageStr) {
					contactInfo.Usage = append(contactInfo.Usage, usageStr)
				}
			}
		}
	}
	if !existing {
		c.Email = append(c.Email, OrganizationConfigContactInfo{Address: address, Usage: usage})
	}
}

func (c *OrganizationConfigContact) AddWebsiteForm(address string, usage []string) {
	address = strings.TrimSpace(address)
	if len(address) == 0 {
		return
	}

	if c.Form == nil {
		c.Form = []OrganizationConfigContactInfo{}
	}
	existing := false
	for _, contactInfo := range c.Form {
		if contactInfo.Address == address {
			existing = true
			for _, usageStr := range usage {
				if !slices.Contains(contactInfo.Usage, usageStr) {
					contactInfo.Usage = append(contactInfo.Usage, usageStr)
				}
			}
		}
	}
	if !existing {
		c.Form = append(c.Form, OrganizationConfigContactInfo{Address: address, Usage: usage})
	}
}

type OrganizationConfigContactInfo struct {
	Address string   `yaml:"address"`
	Usage   []string `yaml:"usage"`
}
