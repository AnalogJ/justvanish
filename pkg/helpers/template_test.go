package helpers

import (
	"github.com/analogj/justvanish/pkg/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTemplatePopulate(t *testing.T) {
	t.Parallel()
	//setup
	orgConfig, err := OrganizationConfig("beenverified.com")
	require.NoError(t, err)
	userConfig := &models.UserConfig{
		FirstName:      "test",
		LastName:       "testLast",
		EmailAddresses: []string{"email@example.com"},
		MailAddresses:  []string{"123 example street, example, CA"},
		PhoneNumbers:   []string{"123-456-7890"},
	}

	//test
	_, err = TemplatePopulate("", "ccpa", "request", userConfig, orgConfig)

	//assert
	require.NoError(t, err)
	//require.Equal(t, "", populatedTemplate)
}

// TODO: add negative tests when data is invalid or incorrect.
//func TestTemplatePopulate_WithMissingData(t *testing.T) {
//	t.Parallel()
//	//setup
//	orgConfig, err := OrganizationConfig("beenverified.com")
//	require.NoError(t, err)
//	userConfig := &models.UserConfig{
//		FirstName:      "test",
//		LastName:       "testLast",
//		EmailAddresses: []string{"email@example.com"},
//		MailAddresses:  []string{"123 example street, example, CA"},
//		//PhoneNumbers:   []string{"123-456-7890"},
//	}
//
//	//test
//	populatedTemplate, err := TemplatePopulate("", "ccpa", "request", userConfig, orgConfig)
//
//	//assert
//	require.NoError(t, err)
//	require.Equal(t, "", populatedTemplate)
//}
