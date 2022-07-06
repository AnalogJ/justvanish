package config

import (
	"errors"
	"github.com/analogj/go-util/utils"
	"github.com/analogj/justvanish/pkg/models"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

// When initializing this class the following methods must be called:
// Config.New
// Config.Init
// This is done automatically when created via the Factory.
type configuration struct {
	*viper.Viper
}

//Viper uses the following precedence order. Each item takes precedence over the item below it:
// explicit call to Set
// flag
// env
// config
// key/value store
// default

func (c *configuration) Init() error {
	c.Viper = viper.New()
	//set defaults

	c.SetDefault("action.dry-run", true)
	c.SetDefault("debug", false)
	c.SetDefault("smtp.hostname", "smtp.gmail.com")
	c.SetDefault("smtp.port", 587)

	//if you want to load a non-standard location system config file (~/drawbridge.yml), use ReadConfig
	c.SetConfigType("yaml")
	c.SetConfigName("config")

	c.SetEnvPrefix("VANISH")
	c.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	c.AutomaticEnv()

	//CLI options will be added via the `Set()` function
	return nil
}

func (c *configuration) ReadConfig(configFilePath string) error {
	//make sure that we specify that this is the correct config path (for eventual WriteConfig() calls)
	c.SetConfigFile(configFilePath)

	configFilePath, err := utils.ExpandPath(configFilePath)
	if err != nil {
		return err
	}

	if !utils.FileExists(configFilePath) {
		log.Printf("No configuration file found at %v. Using Defaults.", configFilePath)
		return errors.New("The configuration file could not be found.")
	}

	log.Printf("Loading configuration file: %s", configFilePath)

	config_data, err := os.Open(configFilePath)
	if err != nil {
		log.Printf("Error reading configuration file: %s", err)
		return err
	}

	err = c.MergeConfig(config_data)
	if err != nil {
		return err
	}

	return nil

}

func (c *configuration) SmtpConfig() *models.SmtpConfig {
	return &models.SmtpConfig{
		Hostname: c.GetString("smtp.hostname"),
		Port:     c.GetInt("smtp.port"),
		Username: c.GetString("smtp.username"),
		Password: c.GetString("smtp.password"),
	}
}
