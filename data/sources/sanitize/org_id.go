package sanitize

import (
	"log"
	"net/url"
	"regexp"
	"strings"
)

//returns a domain name (without www prefix) or empty string
// MUST have http or https prefix
func generateOrgIdFromSanitizedWebsite(sanitizedWebsite string) string {
	if len(sanitizedWebsite) == 0 {
		return ""
	}
	//sanitized website is valid, lets generate an org ID from it
	u, _ := url.Parse(sanitizedWebsite)
	sanitizedOrgId := strings.TrimPrefix(u.Host, "www.")
	return sanitizedOrgId
}

// returns domain name (from email) without www prefix, or empty string
func generateOrgIdFromSanitizedEmail(emailAddress string) string {
	if len(emailAddress) == 0 {
		return ""
	}
	emailParts := strings.Split(emailAddress, "@")
	if strings.Contains(emailParts[1], "gmail.com") {
		return ""
	} else {
		sanitizedWebsite := emailParts[1]
		if !strings.HasPrefix(sanitizedWebsite, "http://") && !strings.HasPrefix(sanitizedWebsite, "https://") {
			sanitizedWebsite = "https://" + sanitizedWebsite
		}
		return generateOrgIdFromSanitizedWebsite(sanitizedWebsite)
	}
}

// returns sanitized organization name as Id
func generateOrgIdFromOrgName(orgName string) string {
	//if the company domain cannot be determined, fallback to the orgname, and sanitize it
	orgName = strings.ToLower(orgName)
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err) //should never happen
	}
	return reg.ReplaceAllString(orgName, "")
}
