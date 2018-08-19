package terse

import (
	"context"

	oidc "github.com/coreos/go-oidc"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"google.golang.org/appengine/urlfetch"
)

type JwtVerifier func(ctx context.Context, rawJwt string) (*Claims, error)

type Claims struct {
	Aud           string `json:"aud"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Exp           int    `json:"exp"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	Iat           int    `json:"iat"`
	Iss           string `json:"iss"`
	Locale        string `json:"locale"`
	Picture       string `json:"picture"`
	Sub           string `json:"sub"`
}

func NewGoogleJwtVerifier(ctx context.Context, clientID string) (JwtVerifier, error) {
	ctx = context.WithValue(ctx, oauth2.HTTPClient, urlfetch.Client(ctx))
	provider, providerErr := oidc.NewProvider(ctx, "https://accounts.google.com")
	if providerErr != nil {
		return nil, errors.Wrap(providerErr, "can't query openid provider")
	}
	verifier := provider.Verifier(&oidc.Config{ClientID: clientID})

	return func(ctx context.Context, rawJwt string) (*Claims, error) {
		idToken, verifyErr := verifier.Verify(ctx, rawJwt)
		if verifyErr != nil {
			return nil, errors.Wrap(verifyErr, "invalid jwt")
		}

		claims := &Claims{}
		claimsErr := idToken.Claims(claims)
		if claimsErr != nil {
			return nil, errors.Wrap(claimsErr, "invalid claims")
		}

		return claims, nil
	}, nil
}
