package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)



type Configuration struct {
	Server struct {
		Port    string
		Logging bool
		GinMode string
		Secret string
		AccessTokenExpireDuration int
		RefreshTokenExpireDuration int
		LimitCountPerRequest int
	}
	App struct {
		Port    string
		Logging bool
		GinMode string
	}
	Broker struct {
		Host     string
		Port     string
		Topic    string
		User     string
		Password string
	}
	Database struct {
		Driver   string
		Name   string
		Logging   bool

	}
}

func CommonConfig() Configuration {
	configFile := "rubix-lib-rest-go-config.json"
	file, err := os.Open(configFile)

	Config := Configuration{}
	if err != nil {
		//app defaults
		Config.Server.Port = "1920"
		Config.Server.AccessTokenExpireDuration = 1
		Config.Server.RefreshTokenExpireDuration = 1
		Config.Server.LimitCountPerRequest = 2
		//app defaults
		//Config.App.Port = "1920"
		Config.App.Logging = false
		Config.App.GinMode = "release"
		//mqtt defaults
		Config.Broker.Host = "0.0.0.0"
		Config.Broker.Port = "1883"
		Config.Broker.Topic = "pub/lora"
		Config.Broker.Port = "1883"
		//db defaults
		Config.Database.Driver = "sqlite"
		Config.Database.Name = "test.db"
		Config.Database.Logging = false
		// generate config file if dont exist
		j, _ := json.Marshal(Config)
		err = ioutil.WriteFile(configFile, j, 0644)
		return Config
	} else {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&Config)
		return Config
	}

}

