package middleware

import (
	"net/http"

	"github.com/gorilla/csrf"
)

func CSRFProtection(next http.Handler) http.Handler {
	return csrf.Protect([]byte("32-byte-long-auth-key"), csrf.Secure(true))(next)
}
