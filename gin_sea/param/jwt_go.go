package param

import (
	"github.com/dgrijalva/jwt-go"
)
type CustomClaims struct {
	ID          int
	Username    string
	AuthorityId string
	BufferTime  int64
	jwt.StandardClaims
}
