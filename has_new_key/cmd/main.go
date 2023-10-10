package main

import (
	"github.com/Flajt/decentproof-backend/has_new_key"
	"github.com/scaleway/serverless-functions-go/local"
)

func main() {
	local.ServeHandler(has_new_key.HandleHasNewKey, local.WithPort(8081))
}
