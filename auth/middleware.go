package auth

import (
	"context"
	"net/http"

	model "github.com/Salaton/screening-test/graph/model"
	db "github.com/Salaton/screening-test/postgres"
)

var DB db.DBClient

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware(db db.DBClient) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			email, err := ParseToken(header)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// //check if customer exists in db
			customer, err := db.GetCustomer(email)

			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &customer)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.Customer {
	raw, _ := ctx.Value(userCtxKey).(*model.Customer)
	return raw
}
