/*
This is a working example of:
- Setting up a Configuration struct
- Creating a function to generate a Config with default values
- Creating a function to load configuration options from a JSON file overriding defaults
- Parsing flags from the command line
*/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

// ConfigOptions are lower level, optional options
type ConfigOptions struct {
	TTL      int    `json:"ttl"`
	Timeout  int    `json:"timeout"`
	Encoding string `json:"encoding"`
}

// Config is the config we are loading
type Config struct {
	Address string        `json:"address"`
	Port    int           `json:"port"`
	Options ConfigOptions `json:"options"`
}

// DefaultConfig creates a new Config loaded with reasonable defaults
func DefaultConfig() *Config {
	return &Config{
		Address: "localhost",
		Port:    8000,
		Options: ConfigOptions{
			TTL:      60,
			Timeout:  90,
			Encoding: "utf8",
		},
	}
}

// NewConfigFromFile loads contents from JSON file at path into a new Config.
// This uses reasonable defaults from DefaultConfig.
func NewConfigFromFile(path string) (*Config, error) {
	config := DefaultConfig()
	data, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	jsonParser := json.NewDecoder(data)
	err = jsonParser.Decode(&config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func main() {
	// Check for a command line flag for the configuration file to load
	configPath := flag.String("config", "/etc/jsonconfigloader/config.json", "Path to JSON configuration file")
	flag.Parse()

	// Load the config
	config, err := NewConfigFromFile(*configPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config)
}
