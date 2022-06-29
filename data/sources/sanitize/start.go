package sanitize

import "fmt"

// Start should try to generate a unique organization id that we can use to compare/lookup organizations
// in general this should always be the website hostname (without scheme eg. http:// or https://), however some
// registry entries have invalid data for the website, so we'll fallback to trying to extract the domain from the email address
// or generate an id from organization name.
func Start(rawWebsite string, rawOrgName string, rawContactEmail string) (finalOrgId string, finalWebsite string, finalContactEmails []string) {

	//sanitize the raw inputs
	sanitizedWebsite, websiteErr := sanitizeWebsite(rawWebsite)
	sanitizedEmailAddresses := sanitizeEmailAddress(rawContactEmail)

	if websiteErr != nil || len(sanitizedWebsite) == 0 {
		//no website extracted, fallback to extracting from email (if possible)
		if len(sanitizedEmailAddresses) > 0 {
			// at least one email address is available, use that
			finalOrgId = generateOrgIdFromSanitizedEmail(sanitizedEmailAddresses[0])
			finalWebsite = finalOrgId
			finalContactEmails = sanitizedEmailAddresses

			if len(finalOrgId) > 0 {
				return
			}
		}

		//fallback
		fmt.Printf("COMPANY NAME FALLBACK -- %s - %s - %s\n", rawWebsite, rawOrgName, rawContactEmail)
		//no email address were extracted (or could not extract valid domain from email), fallback to orgId from name
		finalOrgId = generateOrgIdFromOrgName(rawOrgName)
		finalWebsite = ""
		finalContactEmails = []string{}
	} else {
		//website was extracted successfully
		finalOrgId = generateOrgIdFromSanitizedWebsite(sanitizedWebsite)
		finalWebsite = sanitizedWebsite
		finalContactEmails = sanitizedEmailAddresses
	}
	return
}
