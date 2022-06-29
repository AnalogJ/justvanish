package sanitize

import (
	"fmt"
	"net/url"
	"strings"
)

//these datasets can be kind of nasty, containing empty strings for the domain, or lists instead of single values. we need to handle this
// in a sane way
func sanitizeWebsite(website string) (string, error) {
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
		website = "https://" + website
	}
	//return website, nil

	_, err := url.Parse(website)
	if err != nil {
		return "", err
	}
	return website, nil

}
