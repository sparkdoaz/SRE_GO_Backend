package main

import (
	"os"
)

func NewVaultConnectAddress() string {

	host := os.Getenv("VAULT_HOST")
	if host == "" {
		host = "http://127.0.0.1"
	}

	port := os.Getenv("VAULT_HOST")
	if port == "" {
		port = "8200"
	}

	return host + ":" + port
}
