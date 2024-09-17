module github.com/Flajt/decentproof-backend/webhook

go 1.22

replace github.com/Flajt/decentproof-backend/originstamp => ../originstamp

replace github.com/Flajt/decentproof-backend/scw_secret_wrapper => ../scw_secret_wrapper

replace github.com/Flajt/decentproof-backend/helper => ../helper

replace github.com/Flajt/decentproof-backend/encryption => ../encryption

require (
	github.com/Flajt/decentproof-backend/encryption v0.0.0-20240321002430-57a80da943b1
	github.com/Flajt/decentproof-backend/originstamp v0.0.0-20240321002430-57a80da943b1
	github.com/Flajt/decentproof-backend/scw_secret_wrapper v0.0.0-20240321002430-57a80da943b1
	github.com/rs/zerolog v1.33.0
	github.com/scaleway/serverless-functions-go v0.1.2
	github.com/xhit/go-simple-mail/v2 v2.16.0
)

require (
	github.com/go-test/deep v1.1.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/scaleway/scaleway-sdk-go v1.0.0-beta.30 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/toorop/go-dkim v0.0.0-20240103092955-90b7d1423f92 // indirect
	go.uber.org/mock v0.4.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
