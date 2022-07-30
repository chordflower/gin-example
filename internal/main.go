package app

import (
	"fmt"
	"net/http"

	"github.com/adrg/xdg"
	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigdotenv"
	"github.com/cristalhq/aconfig/aconfigtoml"
	"github.com/cristalhq/aconfig/aconfigyaml"
	"github.com/gin-gonic/gin"
)

// Run the server
func Run() {

	// - Get the global configuration path
	gjson, err := xdg.ConfigFile("example/config.json")
	if err != nil {
		fmt.Printf("Error while obtaining the global configuration file path, %s", err.Error())
		return
	}
	gyaml, err := xdg.ConfigFile("example/config.yml")
	if err != nil {
		fmt.Printf("Error while obtaining the global configuration file path, %s", err.Error())
		return
	}
	gtoml, err := xdg.ConfigFile("example/config.toml")
	if err != nil {
		fmt.Printf("Error while obtaining the global configuration file path, %s", err.Error())
		return
	}
	genv, err := xdg.ConfigFile("example/config.env")
	if err != nil {
		fmt.Printf("Error while obtaining the global configuration file path, %s", err.Error())
		return
	}

	// - Read the configuration
	serverConfiguration := NewWithDefaults()

	loader := aconfig.LoaderFor(serverConfiguration, aconfig.Config{
		SkipFlags:          true,
		EnvPrefix:          "APP_",
		AllowDuplicates:    true,
		AllowUnknownFields: true,
		AllowUnknownEnvs:   true,
		MergeFiles:         false,
		Files: []string{
			gjson, gyaml, gtoml, genv,
			"./config.json", "./config.yml", "./config.toml", "./config.env"},
		FileDecoders: map[string]aconfig.FileDecoder{
			".yml":  aconfigyaml.New(),
			".toml": aconfigtoml.New(),
			".env":  aconfigdotenv.New(),
		},
	})

	if err := loader.Load(); err != nil {
		fmt.Printf("Error while parsing the configuration: %s", err.Error())
		return
	}

	if err := serverConfiguration.Validate(); err != nil {
		fmt.Printf("Error while validating the configuration: %s", err.Error())
		return
	}

	// - Configure the logger

	// - Configure the signals

	// - Configure the database

	// - Configure redis

	// - Configure gin
	r := gin.New()

	// -- Add middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// -- Configure the routes

	// - Start the server
	s := &http.Server{
		Addr:              "127.0.0.1:20034",
		Handler:           r,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		ErrorLog:          nil,
	}

	if err := s.ListenAndServe(); err != nil {
		fmt.Printf("Error while running the server: %s", err.Error())
	}

}
