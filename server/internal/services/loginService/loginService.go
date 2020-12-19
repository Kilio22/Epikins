package loginService

import (
	"epikins-api/internal/services/utils"
	"errors"
	"log"

	"epikins-api/config"

	"github.com/dgrijalva/jwt-go"
)

var IssuerURL = "https://login.microsoftonline.com/" + utils.GetEnvVariable("TENANT_ID") + "/v2.0"

type MyClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func LoginService(appId string, accessToken string) (string, error) {
	if accessToken == "" {
		return "", errors.New("token not found")
	}

	token, err := jwt.ParseWithClaims(accessToken, &MyClaims{}, getKey)
	if err != nil {
		log.Println(err)
		return "", errors.New("bad token")
	}
	if !token.Valid {
		return "", errors.New("bad token")
	}

	claims, ok := token.Claims.(*MyClaims)
	if !ok || !claims.VerifyIssuer(IssuerURL, true) || !claims.VerifyAudience(appId, true) {
		return "", errors.New("bad token")
	}

	if _, ok := config.AuthorizedUsers[claims.Email]; !ok {
		return "", errors.New("you're not authorized to access to this API")
	}
	return claims.Email, nil
}
