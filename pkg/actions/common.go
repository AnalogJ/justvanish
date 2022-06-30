package actions

import (
	"github.com/analogj/justvanish/pkg/config"
	"github.com/analogj/justvanish/pkg/helpers"
	"github.com/analogj/justvanish/pkg/models"
	"github.com/sirupsen/logrus"
)

type CommonAction struct {
	Logger        *logrus.Entry
	Configuration config.Interface
	ActionType    string
}

func (a *CommonAction) Start() error {
	orgList, err := helpers.OrganizationList(&helpers.OrganizationListFilter{
		OrganizationType: a.Configuration.GetString("action.org-type"),
		OrganizationId:   a.Configuration.GetString("action.org-id"),
		RegulationType:   a.Configuration.GetString("action.regulation-type"),
	})
	if err != nil {
		return err
	}

	// get user Configuration (from config file)
	var userConfig models.UserConfig
	err = a.Configuration.UnmarshalKey("user", &userConfig)
	if err != nil {
		return err
	}

	var smtpConfig *models.SmtpConfig
	// get smtp Configuration (from config file)
	if a.Configuration.GetBool("action.dry-run") {
		//TODO: this should come from user specified CONFIG, for now we're going to create a test account
		//smtpConfig := a.Configuration.SmtpConfig()
		smtpConfig, err = helpers.EmailTestSmtpConfig()
		if err != nil {
			return err
		}
	} else {
		smtpConfig, err = helpers.EmailTestSmtpConfig()
		if err != nil {
			return err
		}
	}

	// find Configuration for each organization
	for _, orgId := range orgList {
		orgConfig, err := helpers.OrganizationConfig(orgId)
		if err != nil {
			return err
		}

		if len(orgConfig.Contact.Email) == 0 {
			a.Logger.Warnf("skipping company (%s), no email contact found", orgConfig.OrganizationName)
			continue
		}

		//generate template
		emailContent, err := helpers.TemplatePopulate(
			orgId,
			a.Configuration.GetString("action.regulation-type"),
			a.ActionType,
			&userConfig,
			orgConfig,
		)
		if err != nil {
			return err
		}

		err = helpers.EmailSend(smtpConfig, &userConfig, orgConfig, a.Configuration.GetString("action.regulation-type"), a.ActionType, emailContent)
		if err != nil {
			return err
		}
	}

	return nil
}
