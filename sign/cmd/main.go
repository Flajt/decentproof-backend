package main

import (
	"github.com/Flajt/decentproof-backend/sign"
	"github.com/scaleway/serverless-functions-go/local"
)

func main() {
	local.ServeHandler(sign.HandleSignature, local.WithPort(8082))
}
