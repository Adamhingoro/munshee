package autho

import (
	"context"
	"log"

	"golang.org/x/oauth2"

	oidc "github.com/coreos/go-oidc"
)

type Authenticator struct {
	Provider *oidc.Provider
	Config   oauth2.Config
	Ctx      context.Context
}

func NewAuthenticator() (*Authenticator, error) {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "https://adamhingoro.eu.auth0.com/")
	if err != nil {
		log.Printf("failed to get provider: %v", err)
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     "rPyxM9VIegTVlECGP0AihbO6IZNVHPvB",
		ClientSecret: "yFJZtj0l91PTt17-qSWk7L_Iu2XyMEzf-TX_x_NAw5qu0RrvbnnGktKEzXu3-XYM",
		RedirectURL:  "http://localhost:8002/callback",
		Endpoint: 	  provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
		Ctx:      ctx,
	}, nil
}
