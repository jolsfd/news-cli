package cmd

import "github.com/spf13/viper"

// Config is the struct for the yaml file.
type Config struct {
	Feeds []NewsFeed `mapstructure:"feed"`
	Port  string     `mapstructure:"port"`
}

// SetConfig set defaults, config name and config type.
func SetConfig() {
	viper.SetDefault(portString, port)
	viper.SetDefault(feedString, defaultFeed)

	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
}
