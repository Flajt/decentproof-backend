module github.com/Flajt/decentproof-backend/helper

go 1.20

replace github.com/Flajt/decentproof-backend/scw_secret_wrapper => ../scw_secret_wrapper

require (
	github.com/Flajt/decentproof-backend/scw_secret_wrapper v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.5.1
	go.uber.org/mock v0.3.0
)

require (
	codeberg.org/gusted/mcaptcha v0.0.0-20220723083913-4f3072e1d570 // indirect
	github.com/scaleway/scaleway-sdk-go v1.0.0-beta.21 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
