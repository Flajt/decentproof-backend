module github.com/Flajt/decentproof-backend/decentproof-cron

go 1.20

replace github.com/Flajt/decentproof-backend/scw_secret_wrapper => ../scw_secret_wrapper

require (
	github.com/Flajt/decentproof-backend/scw_secret_wrapper v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.5.1
)

require (
	github.com/scaleway/scaleway-sdk-go v1.0.0-beta.21 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
