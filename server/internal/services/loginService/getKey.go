package loginService

import (
	"encoding/json"
	"epikins-api/internal/services/util"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var EpitechJWKSUri = "https://login.microsoftonline.com/" + util.GetEnvVariable("TENANT_ID") + "/discovery/v2.0/keys"

type JwksKey struct {
	Kid string   `json:"kid"`
	X5c []string `json:"x5c"`
}

type JwksKeys struct {
	Keys []JwksKey `json:"keys"`
}

func getKey(token *jwt.Token) (interface{}, error) {
	res, err := http.DefaultClient.Get(EpitechJWKSUri)
	if err != nil {
		return nil, err
	}

	var jwksKeys JwksKeys
	err = json.NewDecoder(res.Body).Decode(&jwksKeys)
	if err != nil {
		return nil, err
	}
	tokenKid, ok := token.Header["kid"]
	if !ok {
		return nil, errors.New("given token has no \"kid\" header")
	}

	tokenKidString, ok := tokenKid.(string)
	if !ok {
		return nil, errors.New("cannot convert kid to string")
	}
	return getPublicKeyByKid(jwksKeys, tokenKidString)
}
