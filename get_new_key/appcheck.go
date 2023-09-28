package get_new_key

import (
	"context"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

type AppcheckWrapper struct {
	App firebase.App
}

func NewAppcheckWrapper() *AppcheckWrapper {
	//Important GOOGLE_ADMIN_SDK_CREDS needs to be set for this to work
	if json, success := os.LookupEnv("GOOGLE_ADMIN_SDK_CREDS"); !success {
		panic("Admin Creds not found in enviroment")
	} else {
		opt := option.WithCredentialsJSON([]byte(json))
		if app, err := firebase.NewApp(context.Background(), nil, opt); err != nil {
			panic(err)
		} else {
			return &AppcheckWrapper{App: *app}
		}
	}
}

func (appcheckWrapper *AppcheckWrapper) CheckApp(token string) (bool, error) {
	if client, err := appcheckWrapper.App.AppCheck(context.Background()); err != nil {
		panic(err)
	} else {
		if _, err := client.VerifyToken(token); err != nil {
			return false, err
		}
		return true, nil
	}
}