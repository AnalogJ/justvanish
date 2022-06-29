package main

import (
	"fmt"
	"github.com/analogj/justvanish/data/sources/process"
	"github.com/analogj/justvanish/pkg/models"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

func main() {

	storage := map[string]models.OrganizationConfig{}

	log.Printf("Start processing California Data Broker Registry")
	duplicateCount, err := process.ProcessCaliforniaDataBrokerRegistry(storage)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Duplicate entries found after California Data Broker: %d", duplicateCount)

	log.Printf("Start processing Privacy Rights ClearinghouseRegistry")
	duplicateCount, err = process.ProcessPrivacyRightsClearinghouseRegistry(storage)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Duplicate entries found after Privacy Rights Clearinghouse: %d", duplicateCount)

	log.Printf("Start processing Virginia Data Broker Registry")
	duplicateCount, err = process.ProcessVirginiaDataBrokerRegistry(storage)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Duplicate entries found after Virginia Data Broker: %d", duplicateCount)

	//after processing every csv file, write content.
	for orgId, orgConfig := range storage {
		//write file
		yamlStrData, err := yaml.Marshal(&orgConfig)
		if err != nil {
			log.Fatal(err)
		}
		err2 := ioutil.WriteFile(fmt.Sprintf("data/organizations/%s.yaml", orgId), yamlStrData, 0666)
		if err2 != nil {
			log.Fatal(err)
		}
	}
}

//helper functions
