package request

import (
	"fmt"
	"github.com/analogj/justvanish/pkg/config"
	"github.com/analogj/justvanish/pkg/helpers"
	"github.com/analogj/justvanish/pkg/models"
	"github.com/sirupsen/logrus"
)

type RequestAction struct {
	logger        *logrus.Entry
	configuration config.Interface
}

func New(logger *logrus.Entry, configuration config.Interface) (RequestAction, error) {

	return RequestAction{
		logger:        logger,
		configuration: configuration,
	}, nil
}

func (a *RequestAction) Start() error {
	orgList, err := helpers.OrganizationList(&helpers.OrganizationListFilter{
		OrganizationType: a.configuration.GetString("action.org-type"),
		OrganizationId:   a.configuration.GetString("action.org-id"),
		RegulationType:   a.configuration.GetString("action.regulation-type"),
	})
	if err != nil {
		return err
	}

	// get user configuration (from config file)
	var userConfig models.UserConfig
	err = a.configuration.UnmarshalKey("user", &userConfig)
	if err != nil {
		return err
	}

	// find configuration for each organization
	for _, orgId := range orgList {
		orgConfig, err := helpers.OrganizationConfig(orgId)
		if err != nil {
			return err
		}

		//generate template

		emailContent, err := helpers.TemplatePopulate(
			orgId,
			a.configuration.GetString("action.regulation-type"),
			"request",
			&userConfig,
			orgConfig,
		)
		fmt.Printf(emailContent)

		//if not dry run, send email
	}

	return nil
}
