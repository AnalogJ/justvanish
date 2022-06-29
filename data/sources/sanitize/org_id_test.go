package sanitize

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateOrgidFromSanitizedWebsite(t *testing.T) {
	t.Parallel()
	//setup
	table := []struct {
		sanitizedWebsite string
		orgId            string
	}{
		{sanitizedWebsite: "", orgId: ""},
		{sanitizedWebsite: "https://jverify.com", orgId: "jverify.com"},
		{sanitizedWebsite: "https://www.jverify.com", orgId: "jverify.com"},
		{sanitizedWebsite: "www.jverify.com", orgId: ""},
	}

	//test
	for _, r := range table {
		t.Run(r.sanitizedWebsite, func(t *testing.T) {
			require.Equal(t, r.orgId, generateOrgIdFromSanitizedWebsite(r.sanitizedWebsite))
		})
	}
}

func TestGenerateOrgIdFromSanitizedEmail(t *testing.T) {
	t.Parallel()
	//setup
	table := []struct {
		sanitizedEmail string
		orgId          string
	}{
		{sanitizedEmail: "", orgId: ""},
		{sanitizedEmail: "test@jverify.com", orgId: "jverify.com"},
		{sanitizedEmail: "test@www.jverify.com", orgId: "jverify.com"},
	}

	//test
	for _, r := range table {
		t.Run(r.sanitizedEmail, func(t *testing.T) {
			require.Equal(t, r.orgId, generateOrgIdFromSanitizedEmail(r.sanitizedEmail))
		})
	}
}
