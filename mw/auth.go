package mw

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/praelatus/backend/models"
)

// used to store our jwt secret key
var secretKey []byte

func init() {
	if _, err := os.Stat("./.jwt_secret.key"); err == nil {
		keyBytes, err := ioutil.ReadFile("./.jwt_secret.key")
		if err != nil {
			fmt.Println("Error reading keyfile", err)
			os.Exit(1)
		}

		secretKey = keyBytes
		return
	}

	b := make([]byte, 256)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error generating key", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile("./.jwt_secret.key", b, 0600)
	if err != nil {
		fmt.Println("Error writing keyfile", err)
		os.Exit(1)
	}

	secretKey = b
}

// getToken will parse the token out of the headers for the given http.Request
func getToken(r *http.Request) string {
	var tokenStr string

	authHeader := r.Header.Get("Authorization")

	if len(authHeader) > 6 &&
		strings.ToUpper(authHeader[0:6]) == "BEARER" {

		tokenStr = authHeader[7:]

	} else if len(authHeader) > 5 &&
		strings.ToUpper(authHeader[0:5]) == "TOKEN" {

		tokenStr = authHeader[6:]
	}

	return tokenStr
}

// validateToken will validate the token with our jwt library and return the
// corresponding user.
func validateToken(token string) *models.User {
	if token == "" {
		return nil
	}

	tkn, err := jwt.Parse(token, func(tkn *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
			msg := fmt.Sprintf("Unexpected signing method: %v", tkn.Header["alg"])
			return nil, errors.New(msg)
		}

		return secretKey, nil
	})

	if err != nil {
		log.Println("Parse error:", err)
		return nil
	}

	if claims, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {
		e := claims.Valid()
		if e != nil {
			log.Println("Invalid claims:", e)
			return nil
		}

		u := &models.User{}

		e = json.Unmarshal([]byte(claims["sub"].(string)), u)
		if e != nil {
			log.Println("Unable to unmarshal subject:", e)
		}

		return u
	}

	fmt.Println("Claims invalid")
	return nil
}

// JWTSignUser will take the user and return a JWT token signed and with that
// user set as the CurrentUser claim
func JWTSignUser(u models.User) (string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(8 * time.Hour).Unix(),
		Issuer:    "praelatus",
		Subject:   u.String(),
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tkn.SignedString(secretKey)
}

// GetUser will get the current user from the given context
func GetUser(ctx context.Context) *models.User {
	if u, ok := ctx.Value(currentUser).(*models.User); ok {
		return u
	}

	return nil
}

type contextKey string

const currentUser contextKey = "currentUser"

// Auth will check if the token for a request is valid and if so will add the
// current user to the http.Request context
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var u *models.User
		tkn := getToken(r)
		fmt.Println("token", tkn)
		if tkn != "" {
			u = validateToken(tkn)
			fmt.Println(u)
		}

		rq := r.WithContext(context.WithValue(r.Context(), currentUser, u))
		next.ServeHTTP(w, rq)
	})
}
