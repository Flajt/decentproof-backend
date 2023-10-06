package main

import (
	"github.com/Flajt/decentproof-backend/get_new_key"
	"github.com/scaleway/serverless-functions-go/local"
)

func main() {
	local.ServeHandler(get_new_key.HandleGetNewKey, local.WithPort(8080))
}
