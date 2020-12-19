package loginService

import (
	"crypto/rsa"
	"errors"

	"github.com/dgrijalva/jwt-go"
)

func getPublicKeyByKid(jwksKeys JwksKeys, tokenKid string) (*rsa.PublicKey, error) {
	for _, key := range jwksKeys.Keys {
		if key.Kid == tokenKid {
			fullKey := "-----BEGIN CERTIFICATE-----\n" + key.X5c[0] + "\n-----END CERTIFICATE-----"
			rsaPublicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(fullKey))
			if err != nil {
				return nil, err
			} else {
				return rsaPublicKey, nil
			}
		}
	}
	return nil, errors.New("cannot get x5c with kid \"" + tokenKid + "\"")
}
