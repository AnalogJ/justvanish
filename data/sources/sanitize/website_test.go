package sanitize

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSanitizeWebsite(t *testing.T) {
	t.Parallel()
	//setup
	table := []struct {
		rawWebsite       string
		expectError      bool
		sanitizedWebsite string
	}{
		{rawWebsite: "", sanitizedWebsite: "", expectError: true},
		{rawWebsite: "123 my address way, VA", sanitizedWebsite: "", expectError: true},
		{rawWebsite: "http://:weirdcolon.com", sanitizedWebsite: "", expectError: true},
		{rawWebsite: "http://semicolonsuffix.com;", sanitizedWebsite: "http://semicolonsuffix.com", expectError: false},
		{rawWebsite: "missingscheme.com", sanitizedWebsite: "https://missingscheme.com", expectError: false},
		{rawWebsite: "hTtP://MixedCase.com", sanitizedWebsite: "http://mixedcase.com", expectError: false},
		{rawWebsite: "email@address.com", sanitizedWebsite: "https://address.com", expectError: false},
	}

	//test
	for _, r := range table {
		t.Run(r.rawWebsite, func(t *testing.T) {
			website, err := sanitizeWebsite(r.rawWebsite)
			if r.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, r.sanitizedWebsite, website)
			}
		})
	}
}
