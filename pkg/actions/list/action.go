package list

import (
	"fmt"
	"github.com/analogj/go-util/utils"
	"github.com/analogj/justvanish/pkg/config"
	"github.com/analogj/justvanish/pkg/helpers"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"strings"
)

type ListAction struct {
	logger        *logrus.Entry
	configuration config.Interface
}

func New(logger *logrus.Entry, configuration config.Interface) (ListAction, error) {

	return ListAction{
		logger:        logger,
		configuration: configuration,
	}, nil
}

type foundOrg struct {
	OrganizationId   string
	OrganizationType []string
	Usage            []string
}

func (a *ListAction) Start() error {

	foundOrganizations := []foundOrg{}

	orgIds, err := helpers.OrganizationList(&helpers.OrganizationListFilter{
		OrganizationType: a.configuration.GetString("action.org-type"),
		OrganizationId:   a.configuration.GetString("action.org-id"),
		RegulationType:   a.configuration.GetString("action.regulation-type"),
	})
	if err != nil {
		return err
	}

	for _, orgId := range orgIds {
		orgConfig, err := helpers.OrganizationConfig(orgId)
		if err != nil {
			return err
		}

		supportedUsage := map[string]interface{}{}

		for _, contactItem := range orgConfig.Contact.Mail {
			for _, contactItemUsage := range contactItem.Usage {
				supportedUsage[contactItemUsage] = true
			}
		}
		for _, contactItem := range orgConfig.Contact.Email {
			for _, contactItemUsage := range contactItem.Usage {
				supportedUsage[contactItemUsage] = true
			}
		}

		foundOrganizations = append(foundOrganizations, foundOrg{
			OrganizationId:   orgId,
			OrganizationType: orgConfig.OrganizationType,
			Usage:            utils.MapKeys(supportedUsage),
		})
	}

	for _, foundOrganization := range foundOrganizations {
		fmt.Printf("%s [%s] - %s\n",
			foundOrganization.OrganizationId,
			color.HiBlueString("%s", strings.Join(foundOrganization.OrganizationType, ",")),
			color.HiRedString("%s", strings.Join(foundOrganization.Usage, ",")),
		)
	}

	return nil
}
