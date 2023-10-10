module github.com/Flajt/decentproof-backend/sign

go 1.20

replace github.com/Flajt/decentproof-backend/helper => ../helper

replace github.com/Flajt/decentproof-backend/scw_secret_wrapper => ../scw_secret_wrapper

require (
	github.com/Flajt/decentproof-backend/helper v0.0.0-00010101000000-000000000000
	github.com/Flajt/decentproof-backend/scw_secret_wrapper v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.5.1
	github.com/rs/zerolog v1.31.0
	github.com/scaleway/serverless-functions-go v0.1.2
)

require (
	github.com/google/uuid v1.3.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/scaleway/scaleway-sdk-go v1.0.0-beta.21 // indirect
	golang.org/x/sys v0.13.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
