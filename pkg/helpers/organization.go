package helpers

import (
	"fmt"
	"github.com/analogj/justvanish/data"
	"github.com/analogj/justvanish/pkg/models"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
	"path"
	"strings"
)

type OrganizationListFilter struct {
	OrganizationType string
	RegulationType   string
	OrganizationId   string
}

func OrganizationList(filter *OrganizationListFilter) ([]string, error) {
	filteredOrganizationIds := []string{}

	if filter != nil && len(filter.OrganizationId) > 0 {
		//an organization id was provided, filter to this exactly
		_, err := data.Organizations.Open(path.Join("organizations", fmt.Sprintf("%s.yaml", filter.OrganizationId)))
		filteredOrganizationIds = append(filteredOrganizationIds, filter.OrganizationId)
		return filteredOrganizationIds, err
	} else {
		// returning (possibly filtered) list of organization ids
		organizationFileEntries, err := data.Organizations.ReadDir("organizations")
		if err != nil {
			return nil, err
		}

		organizationTypeFilter := ""
		regulationTypeFilter := ""
		if filter != nil && len(filter.OrganizationType) > 0 {
			organizationTypeFilter = filter.OrganizationType
		}
		if filter != nil && len(filter.RegulationType) > 0 {
			regulationTypeFilter = filter.RegulationType
		}

		for _, fileEntry := range organizationFileEntries {
			orgId := strings.TrimSuffix(fileEntry.Name(), ".yaml")

			if len(organizationTypeFilter) == 0 && len(regulationTypeFilter) == 0 {
				// no filters specified, just return all organizations
				filteredOrganizationIds = append(filteredOrganizationIds, orgId)
			} else {
				// 1 or more filters specified
				orgConfig, err := OrganizationConfig(orgId)
				if err != nil {
					return nil, err
				}

				if len(organizationTypeFilter) > 0 && len(regulationTypeFilter) > 0 {
					if slices.Contains(orgConfig.OrganizationType, organizationTypeFilter) && slices.Contains(orgConfig.Regulation, regulationTypeFilter) {
						filteredOrganizationIds = append(filteredOrganizationIds, orgId)
					}
				} else if len(organizationTypeFilter) > 0 && slices.Contains(orgConfig.OrganizationType, organizationTypeFilter) {
					filteredOrganizationIds = append(filteredOrganizationIds, orgId)
				} else if len(regulationTypeFilter) > 0 && slices.Contains(orgConfig.Regulation, regulationTypeFilter) {
					filteredOrganizationIds = append(filteredOrganizationIds, orgId)
				}
			}
		}
	}

	return filteredOrganizationIds, nil
}

func OrganizationConfig(organizationId string) (*models.OrganizationConfig, error) {
	if !strings.HasSuffix(organizationId, ".yaml") {
		organizationId += ".yaml"
	}

	orgYamlConfig, err := data.Organizations.ReadFile(fmt.Sprintf("organizations/%s", organizationId))
	if err != nil {
		return nil, err
	}

	var orgConfig models.OrganizationConfig

	err = yaml.Unmarshal(orgYamlConfig, &orgConfig)
	if err != nil {
		return nil, err
	}
	return &orgConfig, err
}
