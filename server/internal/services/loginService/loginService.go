package loginService

import (
	"context"
	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var IssuerURL = "https://login.microsoftonline.com/" + util.GetEnvVariable("TENANT_ID") + "/v2.0"

type MyClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func isAuthorized(email string, collection *mongo.Collection) (bool, error) {
	res := collection.FindOne(context.TODO(), bson.M{"email": email})

	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, res.Err()
	}
	return true, nil
}

func isEpitechMember(email string) bool {
	return strings.HasSuffix(email, "@epitech.eu")
}

func LoginService(appData *internal.AppData, accessToken string) (string, error) {
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
	if !ok || !claims.VerifyIssuer(IssuerURL, true) || !claims.VerifyAudience(appData.AppId, true) {
		return "", errors.New("bad token")
	}

	ok, err = isAuthorized(claims.Email, appData.UsersCollection)
	if err != nil {
		return "", errors.New("something went wrong: " + err.Error())
	}
	if !ok && !isEpitechMember(claims.Email) {
		return "", errors.New("you're not authorized to access to this API")
	}
	return claims.Email, nil
}
