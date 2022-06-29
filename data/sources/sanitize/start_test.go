package sanitize

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStart(t *testing.T) {
	t.Parallel()
	//setup
	table := []struct {
		rawWebsite         string
		rawOrgName         string
		rawContactEmail    string
		finalOrgId         string
		finalWebsite       string
		finalContactEmails []string
	}{
		{rawWebsite: "", rawOrgName: "Jverify, Inc.", rawContactEmail: "admin@jverify.com",
			finalOrgId: "jverify.com", finalWebsite: "jverify.com", finalContactEmails: []string{"admin@jverify.com"},
		},
	}

	//test
	for _, r := range table {
		t.Run(r.rawWebsite+r.rawOrgName+r.rawContactEmail, func(t *testing.T) {
			actualOrgId, actualWebsite, actualContactEmail := Start(r.rawWebsite, r.rawOrgName, r.rawContactEmail)
			require.Equal(t, r.finalOrgId, actualOrgId)
			require.Equal(t, r.finalWebsite, actualWebsite)
			require.Equal(t, r.finalContactEmails, actualContactEmail)
		})
	}
}
