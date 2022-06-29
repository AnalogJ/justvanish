package process

import (
	"github.com/analogj/justvanish/data/sources/sanitize"
	"github.com/analogj/justvanish/pkg/models"
	"github.com/gocarina/gocsv"
	"os"
	"strings"
)

type PrivacyRightsClearinghouseRegistry struct {
	//"Company Name","Privacy Policy","Data Broker Email",Location,"Data Broker Allows Opt-Out?"
	OrganizatioName     string `csv:"Company Name"`
	ContactEmailAddress string `csv:"Data Broker Email"`
	Website             string `csv:"Privacy Policy"`
	ContactMailAddress  string `csv:"Location"`
	Note                string `csv:"Data Broker Allows Opt-Out?"`
}

func ProcessPrivacyRightsClearinghouseRegistry(storage map[string]models.OrganizationConfig) (int, error) {
	duplicateCount := 0
	// process privacy-rights-clearinghouse.csv
	clientsFile, err := os.OpenFile("data/sources/privacy-rights-clearinghouse.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return duplicateCount, err
	}
	defer clientsFile.Close()
	privacyRightsClearinghouseRegistryRows := []*PrivacyRightsClearinghouseRegistry{}
	if err := gocsv.UnmarshalFile(clientsFile, &privacyRightsClearinghouseRegistryRows); err != nil { // Load clients from file
		return duplicateCount, err
	}

	for _, row := range privacyRightsClearinghouseRegistryRows {

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
		notes = sanitize.AppendNewNote(notes, "Allows Opt-Out? "+row.Note)
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
