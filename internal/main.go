package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Run the server
func Run() {

	// - Get the global configuration path

	// - Read the configuration

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
		ErrorLog:          nil,
	}

	if err := s.ListenAndServe(); err != nil {
		fmt.Printf("Error while running the server: %s", err.Error())
	}

}
