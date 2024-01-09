package middleware

import(
	"net/http"
	"log"
	"context"
	"github.com/aswinbennyofficial/jwt-auth-golang/internal/controllers"
)

// func Authorize(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Parse and validate JWT from request
// 		claims, err := controllers.ParseAndValidateJWT(r)
// 		if err != nil {
// 			log.Println("ERROR WHILE PARSING/VALIDATING JWT: ", err)
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}

// 		// Add claims to the request context for use in subsequent handlers
// 		ctx := r.Context()
// 		ctx = context.WithValue(ctx, "claims", claims)
// 		r = r.WithContext(ctx)

// 		// Call the next handler in the chain
// 		next.ServeHTTP(w, r)
// 	})
// }


func LoginRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse and validate JWT from request
		claims, err := controllers.ParseAndValidateJWT(r)
		if err != nil {
			log.Println("ERROR WHILE PARSING/VALIDATING JWT: ", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Add claims to the request context for use in subsequent handlers
		ctx := r.Context()
		ctx = context.WithValue(ctx, "claims", claims)
		r = r.WithContext(ctx)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}