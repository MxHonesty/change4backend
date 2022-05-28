package authentication

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/MxHonesty/change4backend/logging"
	"github.com/google/uuid"
	"github.com/shaj13/go-guardian/v2/auth"
	"github.com/shaj13/go-guardian/v2/auth/strategies/basic"
	"github.com/shaj13/go-guardian/v2/auth/strategies/token"
	"github.com/shaj13/go-guardian/v2/auth/strategies/union"
	"github.com/shaj13/libcache"
	_ "github.com/shaj13/libcache/fifo"
)

var strategy union.Union
var tokenStrategy auth.Strategy
var cacheObj libcache.Cache

func CreateToken(w http.ResponseWriter, r *http.Request) {
	token := uuid.New().String()
	user := auth.User(r)
	auth.Append(tokenStrategy, token, user)
	body := fmt.Sprintf("token: %s", token)
	w.Write([]byte(body))
}

func SetupGoGuardian() {
	cacheObj = libcache.FIFO.New(0)
	cacheObj.SetTTL(time.Minute * 5)
	cacheObj.RegisterOnExpired(func(key, _ interface{}) {
		cacheObj.Peek(key)
	})
	basicStrategy := basic.NewCached(validateUser, cacheObj)
	tokenStrategy = token.New(token.NoOpAuthenticate, cacheObj)
	strategy = union.New(tokenStrategy, basicStrategy)
}

func validateUser(ctx context.Context, r *http.Request, userName, password string) (auth.Info, error) {
	// here connect to db or any other service to fetch user and validate it.

	if userName == "root" && password == "root" {
		return auth.NewDefaultUser("root", "1", nil, nil), nil
	}

	return nil, fmt.Errorf("invalid credentials")
}

func Middleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logging.InfoLogger.Println("Executing Auth Middleware")
		_, user, err := strategy.AuthenticateRequest(r)
		if err != nil {
			fmt.Println(err)
			code := http.StatusUnauthorized
			http.Error(w, http.StatusText(code), code)
			return
		}
		logging.InfoLogger.Printf("User %s Authenticated\n", user.GetUserName())
		r = auth.RequestWithUser(user, r)
		next.ServeHTTP(w, r)
	})
}
