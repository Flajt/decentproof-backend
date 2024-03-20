module github.com/Flajt/decentproof-backend/sign

go 1.20

replace github.com/Flajt/decentproof-backend/helper => ../helper

replace github.com/Flajt/decentproof-backend/scw_secret_wrapper => ../scw_secret_wrapper

replace github.com/Flajt/decentproof-backend/originstamp => ../originstamp

replace github.com/Flajt/decentproof-backend/encryption => ../encryption

require (
	github.com/Flajt/decentproof-backend/encryption v0.0.0-20240316121540-e651010b448c
	github.com/Flajt/decentproof-backend/helper v0.0.0-20240316121540-e651010b448c
	github.com/Flajt/decentproof-backend/originstamp v0.0.0-20240316121540-e651010b448c
	github.com/Flajt/decentproof-backend/scw_secret_wrapper v0.0.0-20240316121540-e651010b448c
	github.com/rs/zerolog v1.32.0
	github.com/scaleway/serverless-functions-go v0.1.2
	go.uber.org/mock v0.4.0
)

require (
	codeberg.org/gusted/mcaptcha v0.0.0-20220723083913-4f3072e1d570 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/scaleway/scaleway-sdk-go v1.0.0-beta.25 // indirect
	golang.org/x/sys v0.18.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
