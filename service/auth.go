package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/config"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/model"
	"time"
)

// CreateToken method
func CreateToken(email string) (*model.TokenDetailsDTO, error) {

	//_config := config_old.GetConfig()
	_config := config.CommonConfig()

	var err error
	td := &model.TokenDetailsDTO{}

	td.AtExpiresTime = time.Now().Add(time.Hour * time.Duration(_config.Server.AccessTokenExpireDuration))
	td.RtExpiresTime = time.Now().Add(time.Hour * time.Duration(_config.Server.RefreshTokenExpireDuration))

	//create access token
	atClaims := jwt.MapClaims{}
	atClaims["email"] = email
	atClaims["user_uuid"] = "user_uuid"
	atClaims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	atClaims["uuid"] = ""
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	td.AccessToken, err = token.SignedString([]byte(_config.Server.Secret))
	if err != nil {
		return nil, err
	}

	//create refresh  token
	rtClaims := jwt.MapClaims{}
	rtClaims["email"] = email
	rtClaims["user_uuid"] = "user_uuid"
	rtClaims["exp"] = time.Now().Add(time.Hour * 96).Unix() //refresh token expire time config read
	rtClaims["uuid"] = ""
	rtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	td.RefreshToken, err = rtoken.SignedString([]byte(_config.Server.Secret))
	if err != nil {
		return nil, err
	}

	generatedPass, err := GenerateSecureKey(16)
	if err != nil {
		return nil, err
	}
	td.TransmissionKey = generatedPass

	return td, nil
}

// TokenValid method
func TokenValid(bearerToken string) (*jwt.Token, error) {
	token, err := verifyToken(bearerToken)
	if err != nil {
		if token != nil {
			return token, err
		}
		return nil, err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, fmt.Errorf("Unauthorized")
	}
	return token, nil
}

//verifyToken verify token
func verifyToken(tokenString string) (*jwt.Token, error) {
	_config := config.CommonConfig()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(_config.Server.Secret), nil
	})
	if err != nil {
		return token, fmt.Errorf("Unauthorized")
	}
	return token, nil
}
