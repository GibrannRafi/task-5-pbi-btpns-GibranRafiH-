package middleware

import (
	"net/http"

	"github.com/GibrannRafi/task-5-pbi-btpns-GibranRafiH-/app/config"
	"github.com/GibrannRafi/task-5-pbi-btpns-GibranRafiH-/app/helper"
	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response := map[string]string{"message": "Login Terlebih Dahulu!!"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
			helper.ResponseJSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Mengambil token value
		tokenString := c.Value

		claims := &config.JWTclaim{}

		// Parsing token jwt
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				// Token Invalid
				response := map[string]string{"message": "Token tidak valid"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return

			case jwt.ValidationErrorExpired:
				// Token Expired
				response := map[string]string{"message": "Token kadaluarsa"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return

			default:
				// Error lain
				response := map[string]string{"message": "Token tidak valid"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
		}

		if !token.Valid {
			response := map[string]string{"message": "Token tidak valid"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		next.ServeHTTP(w, r)
	})
}
