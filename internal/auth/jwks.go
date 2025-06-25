package auth

import "github.com/MicahParks/keyfunc"

func LoadJWKS() (*keyfunc.JWKS, error) {
	jwksURL := "https://dev-01t73ydzf4mmhno2.us.auth0.com/.well-known/jwks.json"
	jwks, err := keyfunc.Get(jwksURL, keyfunc.Options{})
	return jwks, err
}
