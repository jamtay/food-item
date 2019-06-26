package config

import (
	"log"
	"github.com/BurntSushi/toml"
)

// Represents database server and credentials
type Config struct {
	Server   string
	Database string
}

// ReadFoodItemConfig and parse the configuration file
func (c *Config) ReadFoodItemConfig() {
	if _, err := toml.DecodeFile("config/food_item_config.toml", &c); err != nil {
		log.Fatal(err)
	}
}