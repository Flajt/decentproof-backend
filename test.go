package main

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func main() {
	fmt.Println("Hello, playground")
	if adminSDKCreds, success := os.LookupEnv("GOOGLE_ADMIN_SDK"); success == false {
		panic("We are fucked")
	} else {
		fmt.Println(adminSDKCreds)
		opt := option.WithCredentialsJSON([]byte(adminSDKCreds))
		if app, err := firebase.NewApp(context.Background(), nil, opt); err != nil {
			panic(err)
		} else {
			if appCheck, err := app.AppCheck(context.Background()); err != nil {
				panic(err)
			} else {
				fmt.Println(appCheck.VerifyToken("1244"))
			}
		}

	}
}
