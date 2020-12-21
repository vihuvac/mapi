package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"../helpers"
	jwt "github.com/dgrijalva/jwt-go"
)

// IsAuthorized middleware verifies the token used to perform the request.
func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if len(authHeader) == 0 {
			helpers.ErrorResponse(w, "Bad Request", "Missing authorization header", http.StatusBadRequest)
			return
		}

		authKey := []byte(os.Getenv("AUTH_KEY"))
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error verifying the method assigned to the auth token.")
			}
			return authKey, nil
		})

		if err != nil {
			message := "There was an error verifying the authorization token: " + err.Error()
			helpers.ErrorResponse(w, "Bad Request", message, http.StatusBadRequest)
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)
		if !(claims["user"] == os.Getenv("AUTH_USER") && token.Valid) {
			helpers.ErrorResponse(w, "Unauthorized", "Not authorized", http.StatusUnauthorized)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
