package main

import (
	verify_hash "github.com/Flajt/decentproof-backend/verify-hash"
	"github.com/scaleway/serverless-functions-go/local"
)

func main() {
	local.ServeHandler(verify_hash.HandleHashVerification, local.WithPort(8083))
}
