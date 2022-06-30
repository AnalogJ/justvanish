package helpers

import (
	"bytes"
	"fmt"
	"github.com/analogj/justvanish/data"
	"github.com/analogj/justvanish/pkg/models"
	"text/template"
	"time"
)

type TemplateData struct {
	User interface{}
	Org  *models.OrganizationConfig
	Date string
}

// TODO: add negative tests when data is invalid or incorrect.
func TemplatePopulate(orgId string, regulationType string, actionType string, userData *models.UserConfig, orgConfig *models.OrganizationConfig) (string, error) {
	templateContent, err := data.Templates.ReadFile(fmt.Sprintf("templates/%s/%s.tmpl.md", actionType, regulationType))

	if err != nil {
		return "", err
	}

	tmplData := TemplateData{
		User: userData,
		Org:  orgConfig,
		Date: time.Now().Format("Jan 2 2006"),
	}

	tmpl, err := template.New(fmt.Sprintf("templates/%s/%s.tmpl.md", actionType, regulationType)).Parse(string(templateContent))
	if err != nil {
		return "", err
	}
	var tmplBuff bytes.Buffer
	err = tmpl.Execute(&tmplBuff, tmplData)
	if err != nil {
		return "", err
	}

	return tmplBuff.String(), nil
}
