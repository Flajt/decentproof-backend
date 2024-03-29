module github.com/Flajt/decentproof-backend/encryption

go 1.20

replace github.com/Flajt/decentproof-backend/scw_secret_wrapper => ../scw_secret_wrapper

replace github.com/Flajt/decentproof-backend/helper => ../helper

require (
	github.com/Flajt/decentproof-backend/helper v0.0.0-00010101000000-000000000000
	github.com/Flajt/decentproof-backend/scw_secret_wrapper v0.0.0-20240316121540-e651010b448c
	go.uber.org/mock v0.4.0
)

require (
	codeberg.org/gusted/mcaptcha v0.0.0-20220723083913-4f3072e1d570 // indirect
	github.com/scaleway/scaleway-sdk-go v1.0.0-beta.25 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
