package config

import (
	"github.com/analogj/justvanish/pkg/models"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

// UnsetEnv unsets all envars having prefix and returns a function
// that restores the env. Any newly added envars having prefix are
// also unset by restore. It is idiomatic to use with a defer.
//
//	defer UnsetEnv("ACME_")()
//
// Note that modifying the env may have unpredictable results when
// tests are run with t.Parallel.
// NOTE: This is quick n' dirty from memory; write some tests for
// this code.
func UnsetEnv(prefix string) (restore func()) {
	before := map[string]string{}

	for _, e := range os.Environ() {
		if !strings.HasPrefix(e, prefix) {
			continue
		}

		parts := strings.SplitN(e, "=", 2)
		before[parts[0]] = parts[1]

		os.Unsetenv(parts[0])
	}

	return func() {
		after := map[string]string{}

		for _, e := range os.Environ() {
			if !strings.HasPrefix(e, prefix) {
				continue
			}

			parts := strings.SplitN(e, "=", 2)
			after[parts[0]] = parts[1]

			// Check if the envar previously existed
			v, ok := before[parts[0]]
			if !ok {
				// This is a newly added envar with prefix, zap it
				os.Unsetenv(parts[0])
				continue
			}

			if parts[1] != v {
				// If the envar value has changed, set it back
				os.Setenv(parts[0], v)
			}
		}

		// Still need to check if there have been any deleted envars
		for k, v := range before {
			if _, ok := after[k]; !ok {
				// k is not present in after, so we set it.
				os.Setenv(k, v)
			}
		}
	}
}

func TestConfiguration_InvalidConfigPath(t *testing.T) {
	t.Parallel()

	//setup
	testConfig, _ := Create()

	//test
	err := testConfig.ReadConfig("does_not_exist.yaml")

	//assert
	require.Error(t, err, "should return an error")
}

func TestConfiguration_ShouldUnmarshalUserConfig(t *testing.T) {
	t.Parallel()

	//setup
	testConfig, _ := Create()

	//test
	var userConfig models.UserConfig
	err := testConfig.UnmarshalKey("user", &userConfig)

	//assert
	require.NoError(t, err)
	require.Equal(t, models.UserConfig{FirstName: "", LastName: "", EmailAddresses: []string(nil), MailAddresses: []string(nil), PhoneNumbers: []string(nil)}, userConfig)
}

func TestConfiguration_ShouldUnmarshalSmtpConfig(t *testing.T) {
	//setup
	defer UnsetEnv("VANISH_")()
	os.Setenv("VANISH_SMTP_USERNAME", "test@example.com")
	os.Setenv("VANISH_SMTP_PASSWORD", "my-secure-password")
	testConfig, _ := Create()

	//test
	smtpConfig := testConfig.SmtpConfig()

	//assert
	require.Equal(t, &models.SmtpConfig{
		Hostname: "smtp.gmail.com",
		Port:     587,
		Username: "test@example.com",
		Password: "my-secure-password",
	}, smtpConfig)
}
