package request

import (
	"github.com/analogj/justvanish/pkg/config"
	"github.com/analogj/justvanish/pkg/helpers"
	"github.com/analogj/justvanish/pkg/models"
	"github.com/sirupsen/logrus"
)

type RequestAction struct {
	logger        *logrus.Entry
	configuration config.Interface
	actionType string
}

func New(logger *logrus.Entry, configuration config.Interface) (RequestAction, error) {

	return RequestAction{
		logger:        logger,
		configuration: configuration,
		actionType: "request",
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

	// get smtp configuration (from config file)
	smtpConfig := a.configuration.SmtpConfig()


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
		if err != nil {
			return err
		}

		//if not dry run, send email
		if !a.configuration.GetBool("debug"){
			err := helpers.EmailSend(smtpConfig, &userConfig, orgConfig, a.configuration.GetString("action.regulation-type"), a.actionType, emailContent)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
