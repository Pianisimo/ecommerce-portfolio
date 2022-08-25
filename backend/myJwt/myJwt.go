package myJwt

import (
	"backend/database"
	"backend/models"
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	privateKeyPath = "keys/app.rsa"
	publicKeyPath  = "keys/app.rsa.pub"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// InitJWT Inits the variables verifyKey and signKey.
// Need a private and public rsa keys in the keys' folder.
func InitJWT() error {
	signBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return err
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return err
	}

	verifyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return err
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return err
	}

	return nil
}

func CreateNewTokens(userId int64, role string) (authTokenString, refreshTokenString, csrfSecret string, err error) {
	csrfSecret, err = models.GenerateCSRFSecret()
	if err != nil {
		return
	}

	refreshTokenString, err = createRefreshTokenString(userId, role, csrfSecret)
	if err != nil {
		return
	}

	authTokenString, err = createAuthTokenString(userId, role, csrfSecret)
	if err != nil {
		return
	}

	return
}

func GetUserFromAuthTokenString(authTokenString string) (models.User, error) {
	authToken, err := jwt.ParseWithClaims(authTokenString, &models.TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})
	if err != nil {
		log.Println("Error parsing authTokenString with verifyKey")
		return models.User{}, err
	}

	// Convert received token to models.TokenClaims as authClaims
	authClaims, ok := authToken.Claims.(*models.TokenClaims)
	if !ok {
		fmt.Println("Could not cast parsed authToken.Claims to models.TokenClaims")
		return models.User{}, err
	}

	user, err := database.GetUserById(authClaims.StandardClaims.Id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func CheckIfValid(authTokenString, refreshTokenString, csrfToken string) (value bool) {
	value = false
	if csrfToken == "" {
		fmt.Println("No CSRF token")
		return
	}

	authToken, err := jwt.ParseWithClaims(authTokenString, &models.TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})
	if err != nil {
		log.Println("Error parsing authTokenString with verifyKey")
		return
	}

	// Convert received token to models.TokenClaims as authClaims
	authClaims, ok := authToken.Claims.(*models.TokenClaims)
	if !ok {
		fmt.Println("Could not cast parsed authToken.Claims to models.TokenClaims")
		return
	}

	// Check if authClaims.Csrf matches csrfToken
	if csrfToken != authClaims.Csrf {
		log.Println("CSRF token doesn't match jwt")
		return
	}

	// check if authToken is still valid.
	if authToken.Valid {
		value = true
		return
	} else {
		ve, ok := err.(*jwt.ValidationError)
		if !ok {
			log.Println("Auth token is not valid")
			if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
				log.Println("Auth token has expired")
				if err != nil {
					return
				}
			} else {
				log.Println("error in auth token")
				return
			}
		} else {
			log.Println("error in auth token")
			return
		}
	}

	return
}

func CheckAndRefreshTokens(authTokenString, refreshTokenString, csrfToken string) (newAuthTokenString,
	newRefreshTokenString, newCsrfSecret string, err error) {

	if csrfToken == "" {
		fmt.Println("No CSRF token")
		err = errors.New("unauthorized, no csrf provided")
		return
	}

	authToken, err := jwt.ParseWithClaims(authTokenString, &models.TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})
	if err != nil {
		log.Println("Error parsing authTokenString with verifyKey")
		return
	}

	// Convert received token to models.TokenClaims as authClaims
	authClaims, ok := authToken.Claims.(*models.TokenClaims)
	if !ok {
		fmt.Println("Could not cast parsed authToken.Claims to models.TokenClaims")
		err = errors.New("authToken.Claims can not be casted to models.TokenClaims")
		return
	}

	// Check if authClaims.Csrf matches csrfToken
	if csrfToken != authClaims.Csrf {
		log.Println("CSRF token doesn't match jwt")
		err = errors.New("unauthorized, csrf token doesn't match")
		return
	}

	// check if authToken is still valid.
	if authToken.Valid {
		log.Println("Auth token is valid")
		newCsrfSecret = authClaims.Csrf
		newRefreshTokenString, err = updateRefreshTokenExp(refreshTokenString)
		newAuthTokenString = authTokenString
		return
	} else {
		ve, ok := err.(*jwt.ValidationError)
		if !ok {
			log.Println("Auth token is not valid")
			if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
				log.Println("Auth token has expired")
				//newAuthTokenString, newCsrfSecret, err = updateAuthTokenString(refreshTokenString, authTokenString)
				if err != nil {
					return
				}

				newRefreshTokenString, err = updateRefreshTokenExp(refreshTokenString)
				if err != nil {
					return
				}

				newRefreshTokenString, err = updateRefreshTokenCsrf(newRefreshTokenString, newCsrfSecret)
				if err != nil {
					return
				}
			} else {
				log.Println("error in auth token")
				err = errors.New("error in auth token")
				return
			}
		} else {
			log.Println("error in auth token")
			err = errors.New("error in auth token")
			return
		}
	}

	err = errors.New("unauthorized")
	return
}

// createAuthTokenString returns an authTokenString that needs the signKey to be unraveled.
// the string contains Csrf, ExpiresAt, Id and Role information.
func createAuthTokenString(id int64, role, csrfSecret string) (authTokenString string, err error) {
	authTokenExp := time.Now().Add(models.AuthTokenValidTime).Unix()
	authClaims := models.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: authTokenExp,
			Id:        strconv.FormatInt(id, 10),
		},
		Role: role,
		Csrf: csrfSecret,
	}

	authJwt := jwt.NewWithClaims(jwt.SigningMethodRS256, authClaims)
	authTokenString, err = authJwt.SignedString(signKey)
	return
}

// createRefreshTokenString returns an refreshTokenString that needs the signKey to be unraveled.
// the string contains Csrf, ExpiresAt, Id and Role information.
func createRefreshTokenString(id int64, role, csrfSecret string) (refreshTokenString string, err error) {
	refreshTokenExp := time.Now().Add(models.RefreshedTokenValidTime).Unix()
	refreshClaims := models.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshTokenExp,
			Id:        strconv.FormatInt(id, 10),
		},
		Role: role,
		Csrf: csrfSecret,
	}

	refreshJwt := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshClaims)
	refreshTokenString, err = refreshJwt.SignedString(signKey)
	return
}

func updateRefreshTokenExp(tokenString string) (newTokenString string, err error) {
	oldToken, err := jwt.ParseWithClaims(tokenString, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err != nil {
		return
	}

	oldClaims, ok := oldToken.Claims.(*models.TokenClaims)
	if !ok {
		err = errors.New("can't parse old claims")
		return
	}

	tokenExp := time.Now().Add(models.RefreshedTokenValidTime).Unix()

	claims := models.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExp,
			Id:        oldClaims.StandardClaims.Id,
		},
		Role: oldClaims.Role,
		Csrf: oldClaims.Csrf,
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	newTokenString, err = newToken.SignedString(signKey)
	return
}

/*func updateAuthTokenString(refreshTokenString string, oldAuthTokenString string) (newAuthTokenString, csrfSecret string, err error) {
	refreshToken, err := jwt.ParseWithClaims(refreshTokenString, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	refreshTokenClaims, ok := refreshToken.Claims.(*models.TokenClaims)
	if !ok {
		err = errors.New("Error reading jwt claims")
		return
	}

	if db.CheckRefreshToken(refreshTokenClaims.StandardClaims.Id) {

		if refreshToken.Valid {

			authToken, _ := jwt.ParseWithClaims(oldAuthTokenString, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
				return verifyKey, nil
			})

			oldAuthTokenClaims, ok := authToken.Claims.(*models.TokenClaims)
			if !ok {
				err = errors.New("Error reading jwt claims")
				return
			}

			csrfSecret, err = models.GenerateCSRFSecret()
			if err != nil {
				return
			}

			newAuthTokenString, err = createAuthTokenString(oldAuthTokenClaims.StandardClaims.Id, oldAuthTokenClaims.Role, csrfSecret)

			return
		} else {
			log.Println("Refresh token has expired!")

			db.DeleteRefreshToken(refreshTokenClaims.StandardClaims.Id)

			err = errors.New("Unauthorized")
			return
		}
	} else {
		log.Println("Refresh token has been revoked!")

		err = errors.New("Unauthorized")
		return
	}
}*/

func RevokeTokens(authTokenString, refreshTokenString string) (revokedAuthTokenString string,
	revokedRefreshTokenString string, err error) {
	authToken, err := jwt.ParseWithClaims(authTokenString, &models.TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})
	if err != nil {
		log.Println("Error parsing authTokenString with verifyKey")
		err = errors.New("error parsing authTokenString with verifyKey")
		return
	}

	// Convert received token to models.TokenClaims as authClaims
	authClaims, ok := authToken.Claims.(*models.TokenClaims)
	if !ok {
		fmt.Println("Could not cast parsed authToken.Claims to models.TokenClaims")
		err = errors.New("error casting parsed authToken.Claims to models.TokenClaims")
		return
	}

	authTokenExp := time.Now().Add(-models.AuthTokenValidTime).Unix()
	authClaims.StandardClaims.ExpiresAt = authTokenExp

	revokedAuthToken := jwt.NewWithClaims(jwt.SigningMethodRS256, authClaims)
	revokedAuthTokenString, err = revokedAuthToken.SignedString(signKey)

	refreshToken, err := jwt.ParseWithClaims(refreshTokenString, &models.TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})
	if err != nil {
		log.Println("Error parsing refreshTokenString with verifyKey")
		err = errors.New("error parsing refreshTokenString with verifyKey")
		return
	}

	// Convert received token to models.TokenClaims as authClaims
	refreshClaims, ok := refreshToken.Claims.(*models.TokenClaims)
	if !ok {
		fmt.Println("Could not cast parsed refreshToken.Claims to models.TokenClaims")
		err = errors.New("error casting parsed refreshToken.Claims to models.TokenClaims")
		return
	}

	refreshClaims.StandardClaims.ExpiresAt = authTokenExp

	revokedRefreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshClaims)
	revokedRefreshTokenString, err = revokedRefreshToken.SignedString(signKey)

	id, err := strconv.ParseInt(authClaims.Id, 10, 64)

	err = database.StoreUserTokens(id, revokedAuthTokenString, revokedRefreshTokenString, authClaims.Csrf)
	return
}

func updateRefreshTokenCsrf(refreshTokenString, csrfSecret string) (newRefreshTokenString string, err error) {
	refreshToken, err := jwt.ParseWithClaims(refreshTokenString, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})

	if err != nil {
		return
	}

	refreshClaims, ok := refreshToken.Claims.(*models.TokenClaims)
	if !ok {
		err = errors.New("casting claims error")
		return
	}

	newRefreshClaims := models.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshClaims.StandardClaims.ExpiresAt,
			Id:        refreshClaims.StandardClaims.Id,
		},
		Role: refreshClaims.Role,
		Csrf: csrfSecret,
	}

	newRefreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, newRefreshClaims)

	newRefreshTokenString, err = newRefreshToken.SignedString(signKey)
	return
}

func GrabUUID(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return "", errors.New("error fetching claims")
	})

	if err != nil {
		return "", errors.New("error fetching claims")
	}

	claims, ok := token.Claims.(*models.TokenClaims)
	if !ok {
		return "", errors.New("error fetching claims")
	}

	return claims.StandardClaims.Id, nil
}
