package main

import (
	decentproof_cronjob "github.com/Flajt/decentproof-backend/cronjob"
	"github.com/scaleway/serverless-functions-go/local"
)

func main() {
	local.ServeHandler(decentproof_cronjob.Handle, local.WithPort(8079))
}
