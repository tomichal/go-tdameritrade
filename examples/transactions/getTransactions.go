package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tomichal/go-tdameritrade"
	"golang.org/x/oauth2"
)

func main() {
	// pass an http client with auth
	token := os.Getenv("TDAMERITRADE_CLIENT_ID")
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}
	refreshToken := os.Getenv("TDAMERITRADE_REFRESH_TOKEN")
	if refreshToken == "" {
		log.Fatal("Unauthorized: No refresh token present")
	}

	conf := oauth2.Config{
		ClientID: token,
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://api.tdameritrade.com/v1/oauth2/token",
		},
		RedirectURL: "https://localhost",
	}

	tkn := &oauth2.Token{
		RefreshToken: refreshToken,
	}

	ctx := context.Background()
	tc := conf.Client(ctx, tkn)

	c, err := tdameritrade.NewClient(tc)
	if err != nil {
		log.Fatal(err)
	}

	txns, _, err := c.TransactionHistory.GetTransactions(ctx, "acountid", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", (*txns)[0].TransactionID)
}
