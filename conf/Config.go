// Package conf provides mapping and parsing configuration
package conf

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const envPrefix = "BM"
const defaultConfigPath = "./"
const defaultConfigFile = "config"
const defaultConfigHome = "$HOME/.config"

// Config the application's config struct
type Config struct {
	Port   int64
	Config string
	Logger LogConfig
	Mongo  MongoConfig
}

// LoadConfig load configuration
func LoadConfig(cmd *cobra.Command) (*Config, error) {

	// From command line
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		return nil, err
	}

	// From environment
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// From config file
	if configFile, _ := cmd.Flags().GetString("config"); configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName(defaultConfigFile)
		viper.AddConfigPath(defaultConfigPath)
		viper.AddConfigPath(defaultConfigHome)
	}

	// NOTE: this will require that you have config file somewhere in the paths specified.
	// It can be reading from JSON, TOML, YAML, HCL, and Java properties files.
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return populateConfig(new(Config))
}
