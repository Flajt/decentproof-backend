module github.com/Flajt/decentproof-backend/helper

go 1.22

replace github.com/Flajt/decentproof-backend/scw_secret_wrapper => ../scw_secret_wrapper

require (
	codeberg.org/gusted/mcaptcha v0.0.0-20220723083913-4f3072e1d570
	github.com/Flajt/decentproof-backend/scw_secret_wrapper v0.0.0-20240921113304-cf17ebdad1ff
	github.com/joho/godotenv v1.5.1
	go.uber.org/mock v0.5.0
)

require (
	github.com/kr/text v0.2.0 // indirect
	github.com/scaleway/scaleway-sdk-go v1.0.0-beta.30 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
