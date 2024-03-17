package main

import (
	"os"
)

func NewConsulConnectAddress() string {
	host := os.Getenv("CONSUL_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	port := os.Getenv("CONSUL_PORT")
	if port == "" {
		port = "8500"
	}

	return host + ":" + port
}
