package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"notes/infra"
	"notes/internal/common"
)

func AuthenticationMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")

		if tokenStr == "" {
			common.SendResponse(w, "Missing bearer token", http.StatusBadRequest)
			return
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				fmt.Printf("Token signing method verification error: %s", ok)
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(infra.Cfg.JwtSigningKey), nil
		})

		if err != nil {
			fmt.Printf("Unable to parse token: %s ", err)
			common.SendResponse(w, "Please verify your token", http.StatusBadRequest)
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			handler.ServeHTTP(w, r)
			return
		} else {
			log.Printf("Claims verification has failed: %s", tokenStr)
			common.SendResponse(w, "Provide a valid token", http.StatusBadRequest)
			return
		}
	})
}
