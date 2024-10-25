// config/session.go
package config

import (
	"github.com/gorilla/sessions"
)

var (
	// Initialize a cookie store with a secret key.
	Store = sessions.NewCookieStore([]byte("your-secret-key")) // Use a secure key in production
)
