package jwt

import (
	"context"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://provider.com/o/oauth2/auth",
			TokenURL: "https://provider.com/o/oauth2/token",
		},
		RedirectURL: "",
		Scopes:      nil,
	}

}
