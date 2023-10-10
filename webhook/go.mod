module github.com/Flajt/decentproof-backend/webhook

go 1.20

replace github.com/Flajt/decentproof-backend/originstamp => ../originstamp

require (
	github.com/Flajt/decentproof-backend/originstamp v0.0.0-00010101000000-000000000000
	github.com/scaleway/serverless-functions-go v0.1.2
	github.com/xhit/go-simple-mail/v2 v2.16.0
)

require (
	github.com/go-test/deep v1.1.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/rs/zerolog v1.31.0 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/toorop/go-dkim v0.0.0-20201103131630-e1cd1a0a5208 // indirect
	golang.org/x/sys v0.13.0 // indirect
)
