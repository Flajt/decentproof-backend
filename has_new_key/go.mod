module github.com/Flajt/decentproof-backend/has_new_key

go 1.22

replace github.com/Flajt/decentproof-backend/helper => ../helper

replace github.com/Flajt/decentproof-backend/scw_secret_wrapper => ../scw_secret_wrapper

require (
	github.com/Flajt/decentproof-backend/helper v0.0.0-20240316121540-e651010b448c
	github.com/Flajt/decentproof-backend/scw_secret_wrapper v0.0.0-20240321002430-57a80da943b1
	github.com/rs/zerolog v1.33.0
	github.com/scaleway/serverless-functions-go v0.1.2
)

require (
	codeberg.org/gusted/mcaptcha v0.0.0-20220723083913-4f3072e1d570 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/scaleway/scaleway-sdk-go v1.0.0-beta.28 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	go.uber.org/mock v0.4.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
