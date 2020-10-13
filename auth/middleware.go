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

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			// splitToken := strings.Split(reqToken, "Bearer ")
			// reqToken = splitToken[1]

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			// tokenStr := header
			// Should return the Username that was used to generate the token
			username, err := ParseToken(header)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			//check if user exists in db
			user, err := DB.GetUser(username)
			// var user model.User
			// row := ps.db.Table("users").Where("username = ?", username).Select("id,username").Row()
			// row.Scan(&user)
			if err != nil {
				next.ServeHTTP(w, r)

				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}
