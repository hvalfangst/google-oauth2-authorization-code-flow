package configuration

import (
	"encoding/json"
	"errors"
	"os"
)

// Configuration represents the JSON structure.
type Configuration struct {
	Db     Db     `json:"db"`
	Google Google `json:"google"`
}

type Db struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Database string `json:"database"`
}

type Google struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func Get(key string) (interface{}, error) {
	file, err := os.Open("src/configuration.json")
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var config Configuration
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	// Determine which key to access and return the corresponding value
	switch key {
	case "db":
		return config.Db, nil
	case "google":
		return config.Google, nil
	default:
		return nil, errors.New("invalid key")
	}
}
