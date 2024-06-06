package middleware

import (
    "net/http"
    "strings"

    "github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("secret_key")

// Middleware untuk memeriksa token JWT
func JwtAuthentication(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing authorization header", http.StatusUnauthorized)
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")

        claims := &jwt.StandardClaims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Add user ID to context if needed
        // ctx := context.WithValue(r.Context(), "userID", claims.Subject)
        // r = r.WithContext(ctx)

        next.ServeHTTP(w, r)
    })
}
