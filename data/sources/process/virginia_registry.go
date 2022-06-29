package process

import (
	"github.com/analogj/justvanish/data/sources/sanitize"
	"github.com/analogj/justvanish/pkg/models"
	"github.com/gocarina/gocsv"
	"os"
	"strings"
)

type VirginiaDataBrokerRegistry struct {
	OrganizatioName     string `csv:"Data Broker Name:"`
	ContactEmailAddress string `csv:"Email Address:"`
	Website             string `csv:"Primary Internet Address:"`
	ContactMailAddress  string `csv:"Address:"`
	Note1               string `csv:"3. Does the data broker permit a consumer to opt out of the data broker’s collection of brokered personal information, opt out ofits databases or opt out of certain sales of data? :"`
	Note2               string `csv:"a. What was the method for requesting an opt-out?"`
	Note3               string `csv:"8. Any additional information or explanation the data broker chooses to provide concerning its data collection practices:"`
}

func ProcessVirginiaDataBrokerRegistry(storage map[string]models.OrganizationConfig) (int, error) {
	duplicateCount := 0
	// process virginia-data-broker-registry.csv
	clientsFile, err := os.OpenFile("data/sources/virginia-data-broker-registry.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return duplicateCount, err
	}
	defer clientsFile.Close()
	virginiaDataBrokerRegistryRows := []*VirginiaDataBrokerRegistry{}
	if err := gocsv.UnmarshalFile(clientsFile, &virginiaDataBrokerRegistryRows); err != nil { // Load clients from file
		return duplicateCount, err
	}

	for _, row := range virginiaDataBrokerRegistryRows {

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
			//panic(errors.New("domain alreay existst "+ domainName))
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

		//fmt.Println("%v", orgConfig)

	}
	return duplicateCount, nil
}
