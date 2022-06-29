package sanitize

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSanitizeEmailAddress(t *testing.T) {
	t.Parallel()
	//setup
	table := []struct {
		rawEmail       string
		sanitizedEmail []string
	}{
		{rawEmail: "protected [at] example.com", sanitizedEmail: []string{"protected@example.com"}},
		{rawEmail: "UpperCase@example.com", sanitizedEmail: []string{"uppercase@example.com"}},
		{rawEmail: " spaces@example.com ", sanitizedEmail: []string{"spaces@example.com"}},
		{rawEmail: "Strip Names <stripnames@example.com>", sanitizedEmail: []string{"stripnames@example.com"}},
		{rawEmail: "Name comma, namecomma@example.com", sanitizedEmail: []string{"namecomma@example.com"}},
		{rawEmail: "Ignore invalid", sanitizedEmail: []string{}},
		{rawEmail: "noAtsymbol", sanitizedEmail: []string{}},
		{rawEmail: "", sanitizedEmail: []string{}},
	}

	//test
	for _, r := range table {
		t.Run(r.rawEmail, func(t *testing.T) {
			require.Equal(t, r.sanitizedEmail, sanitizeEmailAddress(r.rawEmail))
		})
	}
}
