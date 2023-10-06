package main

import (
	"github.com/Flajt/decentproof-backend/webhook"
	"github.com/scaleway/serverless-functions-go/local"
)

func main() {
	local.ServeHandler(webhook.HandleWebhookCallBack, local.WithPort(8084))
}
