package sanitize

import "strings"

func sanitizeEmailAddress(rawEmailAddress string) []string {
	sanitizedEmailAddresses := []string{}
	for _, emailAddress := range strings.Split(strings.Replace(strings.ToLower(rawEmailAddress), " [at] ", "@", -1), " ") {
		emailAddress = strings.Trim(emailAddress, ",;|><:")
		if strings.Contains(emailAddress, "@") {
			sanitizedEmailAddresses = append(sanitizedEmailAddresses, emailAddress)
		}
	}
	return sanitizedEmailAddresses
}
