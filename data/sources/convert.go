package main

import (
	"fmt"
	"github.com/analogj/justvanish/pkg/models"
	"github.com/gocarina/gocsv"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net/url"
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

func main() {
	// process california-data-broker-registry.csv
	clientsFile, err := os.OpenFile("data/sources/california-data-broker-registry.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()
	californiaDataBrokerRegistryRows := []*CaliforniaDataBrokerRegistry{}
	if err := gocsv.UnmarshalFile(clientsFile, &californiaDataBrokerRegistryRows); err != nil { // Load clients from file
		panic(err)
	}

	storage := map[string]models.OrganizationConfig{}

	duplicateCount := 0
	for _, row := range californiaDataBrokerRegistryRows {

		u, err := url.Parse(row.Website)
		if err != nil {
			panic(err)
		}

		domainName := strings.TrimPrefix(strings.ToLower(u.Host), "www.")

		orgConfig, existing := storage[domainName]
		if !existing {
			orgConfig = models.OrganizationConfig{
				OrganizationName: row.OrganizatioName,
				Website:          strings.ToLower(row.Website),
				OrganizationType: []string{"databroker"},
				Regulation:       []string{"ccpa"},
				Contact:          models.OrganizationConfigContact{},
			}
		} else {
			duplicateCount += 1
			//panic(errors.New("domain alreay existst "+ domainName))
		}

		//safe updates
		notes := orgConfig.Notes
		if notes != row.Note1 {
			notes += "\n" + row.Note1
		}
		if notes != row.Note2 {
			notes += "\n" + row.Note2
		}
		if notes != row.Note3 {
			notes += "\n" + row.Note3
		}
		orgConfig.Notes = notes
		orgConfig.Contact.AddEmail(strings.ToLower(strings.Replace(row.ContactEmailAddress, " [at] ", "@", -1)), []string{"request.ccpa", "delete.ccpa", "donotsell.ccpa"})
		orgConfig.Contact.AddMail(row.ContactMailAddress, []string{"request.ccpa", "delete.ccpa", "donotsell.ccpa"})

		//last
		storage[domainName] = orgConfig

		fmt.Println("%v", orgConfig)

		//write file
		yamlStrData, err := yaml.Marshal(&orgConfig)
		if err != nil {
			log.Fatal(err)
		}
		err2 := ioutil.WriteFile(fmt.Sprintf("data/organizations/%s.yaml", domainName), yamlStrData, 0666)
		if err2 != nil {
			log.Fatal(err2)
		}
	}
	log.Printf("Duplicate entries found: %d", duplicateCount)
}
