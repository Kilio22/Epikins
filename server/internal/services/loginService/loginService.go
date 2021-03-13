package loginService

import (
	"context"
	"log"
	"net/http"
	"strings"

	"epikins-api/config"
	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/dgrijalva/jwt-go"
)

var IssuerURL = "https://login.microsoftonline.com/" + util.GetEnvVariable(config.TenantIdKey) + "/v2.0"

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

func LoginService(appData *internal.AppData, accessToken string) (string, internal.MyError) {
	if accessToken == "" {
		return "", util.GetMyError("token not found", nil, http.StatusUnauthorized)
	}

	token, err := jwt.ParseWithClaims(accessToken, &MyClaims{}, getKey)
	if err != nil {
		log.Println(err)
		return "", util.GetMyError("bad token", nil, http.StatusUnauthorized)
	}
	if !token.Valid {
		return "", util.GetMyError("bad token", nil, http.StatusUnauthorized)
	}

	claims, ok := token.Claims.(*MyClaims)
	if !ok || !claims.VerifyIssuer(IssuerURL, true) || !claims.VerifyAudience(appData.AppId, true) {
		return "", util.GetMyError("bad token", nil, http.StatusUnauthorized)
	}

	ok, err = isAuthorized(claims.Email, appData.UsersCollection)
	if err != nil {
		return "", util.GetMyError("can't log user in", err, http.StatusInternalServerError)
	}
	if !ok && !isEpitechMember(claims.Email) {
		return "", util.GetMyError("you're not authorized to access to this API", err, http.StatusForbidden)
	}
	return claims.Email, internal.MyError{}
}
