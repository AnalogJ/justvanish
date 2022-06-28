package main

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
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
	for _, row := range californiaDataBrokerRegistryRows {
		fmt.Println("%v", row)
	}

}
