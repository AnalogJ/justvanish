package process

import (
	"github.com/analogj/justvanish/data/sources/sanitize"
	"github.com/analogj/justvanish/pkg/models"
	"github.com/gocarina/gocsv"
	"os"
	"strings"
)

type CaliforniaDataBrokerRegistry struct {
	OrganizatioName     string `csv:"Data Broker Name"`
	ContactEmailAddress string `csv:"Email Address"`
	Website             string `csv:"Website URL"`
	ContactMailAddress  string `csv:"Physical Address"`
	Note1               string `csv:"How a consumer may opt out of sale or submit requests under the CCPA"`
	Note2               string `csv:"How a protected individual can demand deletion of information posted online under Gov. Code sections 6208.1(b) or 6254.21(c)(1)"`
	Note3               string `csv:"Additional information about data collecting practices"`
	NotUsed             string `csv:"Date Added"`
}

func ProcessCaliforniaDataBrokerRegistry(storage map[string]models.OrganizationConfig) (int, error) {
	duplicateCount := 0
	// process california-data-broker-registry.csv
	clientsFile, err := os.OpenFile("data/sources/california-data-broker-registry.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return duplicateCount, err
	}
	defer clientsFile.Close()
	californiaDataBrokerRegistryRows := []*CaliforniaDataBrokerRegistry{}
	if err := gocsv.UnmarshalFile(clientsFile, &californiaDataBrokerRegistryRows); err != nil { // Load clients from file
		return duplicateCount, err
	}

	for _, row := range californiaDataBrokerRegistryRows {

		orgId, website, contactEmails := sanitize.Start(row.Website, row.OrganizatioName, row.ContactEmailAddress)

		orgConfig, existing := storage[orgId]
		if !existing {
			orgConfig = models.OrganizationConfig{
				OrganizationName: row.OrganizatioName,
				Website:          website,
				OrganizationType: []string{"databroker"},
				Regulation:       []string{"ccpa"},
				Contact:          models.OrganizationConfigContact{},
			}
		} else {
			duplicateCount += 1
		}

		//safe updates
		notes := []string{strings.TrimSpace(orgConfig.Notes)}
		notes = sanitize.AppendNewNote(notes, row.Note1, row.Note2, row.Note3)
		orgConfig.Notes = strings.Join(notes, "\n")
		for _, emailAddress := range contactEmails {
			orgConfig.Contact.AddEmail(emailAddress, []string{"request.ccpa", "delete.ccpa", "donotsell.ccpa"})
		}
		orgConfig.Contact.AddMail(row.ContactMailAddress, []string{"request.ccpa", "delete.ccpa", "donotsell.ccpa"})

		//last
		storage[orgId] = orgConfig

	}
	return duplicateCount, nil
}
