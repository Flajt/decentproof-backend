package helper

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"codeberg.org/gusted/mcaptcha"
)

func Authenticate(req *http.Request, apiKeys []string, verify bool) (bool, error) {
	if authHeader := req.Header.Get("Authorization"); authHeader != "" {
		apiKey := strings.Split(authHeader, " ")[1]
		match := false
		for _, key := range apiKeys {
			if apiKey == key {
				match = true
				break
			}
		}
		return match, nil
	} else if authHeader := req.Header.Get("X-MCAPTCHA-TOKEN"); authHeader != "" && verify {
		success, err := mcaptcha.Verify(context.Background(), &mcaptcha.VerifyOpts{Token: authHeader, Secret: os.Getenv("MCAPTCHA_SECRET"), Sitekey: os.Getenv("MCAPTCHA_SITEKEY"), InstanceURL: os.Getenv("MCAPTCHA_INSTANCE_URL")})
		if err != nil {
			return false, err
		} else {
			return success, nil
		}
	} else {
		return false, fmt.Errorf("no auth header found")
	}
}
