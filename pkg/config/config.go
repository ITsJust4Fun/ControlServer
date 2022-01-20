package config

import (
	"encoding/json"
	"os"
	"time"
)

type Config struct {
	Port                string
	DatabaseURI         string
	DatabaseName        string
	DatabaseTimeout     time.Duration
	TokenExpires        time.Duration
	RefreshTokenExpires time.Duration
	DeviceTokenExpires  time.Duration
	SecretKey           []byte
	IsLoaded            bool
}

type JsonData struct {
	Port                string `json:"port"`
	DatabaseURI         string `json:"database_uri"`
	DatabaseName        string `json:"database_name"`
	DatabaseTimeout     int    `json:"database_timeout"`
	TokenExpires        int    `json:"token_expires"`
	RefreshTokenExpires int    `json:"refresh_token_expires"`
	DeviceTokenExpires  int    `json:"device_token_expires"`
	SecretKey           string `json:"secret_key"`
}

var defaultConfig = Config{
	Port:                "8081",
	DatabaseURI:         "mongodb://127.0.0.1:27017",
	DatabaseName:        "control_server",
	DatabaseTimeout:     10 * time.Second,
	TokenExpires:        15 * time.Minute,
	RefreshTokenExpires: 24 * time.Hour,
	DeviceTokenExpires:  24 * time.Hour,
	SecretKey:           []byte("secret_key"),
	IsLoaded:            false,
}

var loadedConfig = Config{IsLoaded: false}

func ReadConfigFile() {
	file, err := os.Open("config.json")

	if err != nil {
		return
	}

	jsonData := new(JsonData)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(jsonData)

	if err != nil {
		return
	}

	loadedConfig.Port = jsonData.Port
	loadedConfig.DatabaseURI = jsonData.DatabaseURI
	loadedConfig.DatabaseName = jsonData.DatabaseName
	loadedConfig.DatabaseTimeout = time.Duration(jsonData.DatabaseTimeout) * time.Second
	loadedConfig.TokenExpires = time.Duration(jsonData.TokenExpires) * time.Minute
	loadedConfig.RefreshTokenExpires = time.Duration(jsonData.RefreshTokenExpires) * time.Hour
	loadedConfig.DeviceTokenExpires = time.Duration(jsonData.DeviceTokenExpires) * time.Hour
	loadedConfig.IsLoaded = true
}

func GetConfig() Config {
	if !loadedConfig.IsLoaded {
		return defaultConfig
	}

	return loadedConfig
}
