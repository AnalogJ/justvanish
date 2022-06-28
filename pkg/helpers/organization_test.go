package helpers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOrganizationList_WithoutFilter(t *testing.T) {
	t.Parallel()
	//setup

	//test
	orgIdList, err := OrganizationList(nil)

	//assert
	require.NoError(t, err)
	require.True(t, len(orgIdList) > 0)
}

func TestOrganizationList_WithOrgIdFilter(t *testing.T) {
	t.Parallel()
	//setup

	//test
	orgIdList, err := OrganizationList(&OrganizationListFilter{OrganizationId: "beenverified.com"})

	//assert
	require.NoError(t, err)
	require.Equal(t, orgIdList, []string{"beenverified.com"})
}

func TestOrganizationList_WithOrgIdFilter_DoesNotExist(t *testing.T) {
	t.Parallel()
	//setup

	//test
	_, err := OrganizationList(&OrganizationListFilter{OrganizationId: "doesnotexistsdfsdfsdf.com"})

	//assert
	require.Error(t, err)
}

func TestOrganizationConfig_WithValidOrgId(t *testing.T) {
	t.Parallel()
	//setup

	//test
	orgConfig, err := OrganizationConfig("beenverified.com")

	//assert
	require.NoError(t, err)
	require.Equal(t, "Beenverified, Inc", orgConfig.OrganizationName)
}

func TestOrganizationConfig_WithValidOrgId_WithYamlSuffix(t *testing.T) {
	t.Parallel()
	//setup

	//test
	orgConfig, err := OrganizationConfig("beenverified.com.yaml")

	//assert
	require.NoError(t, err)
	require.Equal(t, "Beenverified, Inc", orgConfig.OrganizationName)
}

func TestOrganizationConfig_WithInvalidOrgId(t *testing.T) {
	t.Parallel()
	//setup

	//test
	_, err := OrganizationConfig("sdfsdfsdfsdfed.com.yaml")

	//assert
	require.Error(t, err)
}

func TestOrganizationConfig_WithEmptyOrgId(t *testing.T) {
	t.Parallel()
	//setup

	//test
	_, err := OrganizationConfig("")

	//assert
	require.Error(t, err)
}
