package main

import (
	"fmt"
	"github.com/analogj/justvanish/pkg/models"
	"github.com/gocarina/gocsv"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func main() {

	storage := map[string]models.OrganizationConfig{}
	log.Printf("Start processing California Data Broker Registry")
	duplicateCount, err := ProcessCaliforniaDataBrokerRegistry(storage)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Duplicate entries found after California Data Broker: %d", duplicateCount)

	log.Printf("Start processing Privacy Rights ClearinghouseRegistry")
	duplicateCount, err = ProcessPrivacyRightsClearinghouseRegistry(storage)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Duplicate entries found after Privacy Rights Clearinghouse: %d", duplicateCount)

	log.Printf("Start processing Virginia Data Broker Registry")
	duplicateCount, err = ProcessVirginiaDataBrokerRegistry(storage)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Duplicate entries found after Virginia Data Broker: %d", duplicateCount)

	//after processing every csv file, write content.
	for domainName, orgConfig := range storage {
		//write file
		yamlStrData, err := yaml.Marshal(&orgConfig)
		if err != nil {
			log.Fatal(err)
		}
		err2 := ioutil.WriteFile(fmt.Sprintf("data/organizations/%s.yaml", domainName), yamlStrData, 0666)
		if err2 != nil {
			log.Fatal(err)
		}
	}
}

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

		domainName, err := SanitizeOrganizationIdDomain(row.Website)
		if err != nil {
			//this data is sanitized pretty well, we should just skip the handful of invalid entries
			//return duplicateCount, err
			fmt.Printf("falling back to sanitized company name as id (%s)\n", row.OrganizatioName)
			domainName = SanitizeOrganizationIdFromNameFallback(row.OrganizatioName)
		}

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
		notes := []string{strings.TrimSpace(orgConfig.Notes)}
		notes = AppendNewNote(notes, row.Note1)
		notes = AppendNewNote(notes, row.Note2)
		notes = AppendNewNote(notes, row.Note3)
		orgConfig.Notes = strings.Join(notes, "\n")
		for _, emailAddress := range SanitizeEmailAddress(row.ContactEmailAddress) {
			orgConfig.Contact.AddEmail(emailAddress, []string{"request.ccpa", "delete.ccpa", "donotsell.ccpa"})
		}
		orgConfig.Contact.AddMail(row.ContactMailAddress, []string{"request.ccpa", "delete.ccpa", "donotsell.ccpa"})

		//last
		storage[domainName] = orgConfig

		//fmt.Println("%v", orgConfig)

	}
	return duplicateCount, nil
}

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

		//this dataset is kind of nasty, the Website column may have multiple entries and sometimes its empty.
		var domainName string
		domainName, err = SanitizeOrganizationIdDomain(row.Website)
		if err != nil {
			//fallback to email address
			emailParts := strings.Split(row.ContactEmailAddress, "@")
			if len(emailParts) == 1 {
				fmt.Printf("skipping company: %s - invalid website invalid email\n", row.OrganizatioName)
				continue
			} else {
				fmt.Printf("falling back to email domain for company %s - %s (%s)\n", row.OrganizatioName, row.ContactEmailAddress, emailParts[1])
				if strings.Contains(emailParts[1], "gmail.com") {
					continue
				}

				domainName, err = SanitizeOrganizationIdDomain(emailParts[1])
			}
		}
		if len(row.Website) == 0 {
			row.Website = domainName
		}

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
		notes := []string{strings.TrimSpace(orgConfig.Notes)}
		notes = AppendNewNote(notes, "Allows Opt-Out? "+row.Note)
		orgConfig.Notes = strings.Join(notes, "\n")
		for _, emailAddress := range SanitizeEmailAddress(row.ContactEmailAddress) {
			orgConfig.Contact.AddEmail(emailAddress, []string{"request.ccpa", "delete.ccpa", "donotsell.ccpa"})
		}
		orgConfig.Contact.AddMail(row.ContactMailAddress, []string{"request.ccpa", "delete.ccpa", "donotsell.ccpa"})

		//last
		storage[domainName] = orgConfig
		//fmt.Println("%v", orgConfig)

	}
	return duplicateCount, nil
}

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

		domainName, err := SanitizeOrganizationIdDomain(row.Website)
		if err != nil {
			fmt.Printf("falling back to sanitized company name as id (%s)\n", row.OrganizatioName)
			domainName = SanitizeOrganizationIdFromNameFallback(row.OrganizatioName)
			//return duplicateCount, err
		}

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
		notes := []string{strings.TrimSpace(orgConfig.Notes)}
		notes = AppendNewNote(notes, row.Note1)
		notes = AppendNewNote(notes, row.Note2)
		notes = AppendNewNote(notes, row.Note3)
		orgConfig.Notes = strings.Join(notes, "\n")
		for _, emailAddress := range SanitizeEmailAddress(row.ContactEmailAddress) {
			orgConfig.Contact.AddEmail(emailAddress, []string{"request.ccpa", "delete.ccpa", "donotsell.ccpa"})
		}
		orgConfig.Contact.AddMail(row.ContactMailAddress, []string{"request.ccpa", "delete.ccpa", "donotsell.ccpa"})

		//last
		storage[domainName] = orgConfig

		//fmt.Println("%v", orgConfig)

	}
	return duplicateCount, nil
}

//helper functions
func AppendNewNote(currentNotes []string, newNote string) []string {
	newNote = strings.TrimSpace(newNote)
	if !slices.Contains(currentNotes, newNote) {
		currentNotes = append(currentNotes, newNote)
	}
	return currentNotes
}

func SanitizeEmailAddress(rawEmailAddress string) []string {
	sanitizedEmailAddresses := []string{}
	for _, emailAddress := range strings.Split(strings.Replace(strings.ToLower(rawEmailAddress), " [at] ", "@", -1), " ") {
		emailAddress = strings.Trim(emailAddress, ",;|><:")
		if strings.Contains(emailAddress, "@") {
			sanitizedEmailAddresses = append(sanitizedEmailAddresses, emailAddress)
		}
	}
	return sanitizedEmailAddresses
}

//these datasets can be kind of nasty, containing empty strings for the domain, or lists instead of single values. we need to handle this
// in a sane way
func SanitizeOrganizationIdDomain(website string) (string, error) {
	website = strings.TrimSpace(strings.ToLower(website))
	//if space delimited list is provided, split and take the first item.
	website = strings.TrimSpace(strings.Split(website, " ")[0])
	website = strings.Trim(website, ",;|><:")
	if strings.Contains(website, "@") {
		website = strings.Split(website, "@")[1]
	}
	if !strings.Contains(website, ".") {
		//this website doesnt contain any "." characters, this is probably invalid
		return "", fmt.Errorf("invalid website provided: %s", website)
	}
	if len(website) == 0 {
		return "", fmt.Errorf("invalid website provided")
	}

	// if a scheme is not provided url.Parse will return an empty string for u.Host
	if !strings.HasPrefix(website, "http") {
		website = "http://" + website
	}

	u, err := url.Parse(website)
	if err != nil {
		return "", err
	}

	//remove www prefix if present.
	return strings.TrimPrefix(u.Host, "www."), nil
}

func SanitizeOrganizationIdFromNameFallback(orgName string) string {
	//if the company domain cannot be determined, fallback to the orgname, and sanitize it
	orgName = strings.ToLower(orgName)
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(orgName, "")
}
