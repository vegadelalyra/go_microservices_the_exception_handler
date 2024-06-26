package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	CompanyId string `json:"comapnyId,omitempty"`
	Username  string `json:"username,omitempty"`
	Roles     []int  `json:"roles,omitempty"`
	jwt.StandardClaims
}

func (claims JwtClaims) Valid() error {
	var now = time.Now().UTC().Unix()
	flags, _ := GetFlags()
	url, _ := flags.GetApplicationUrl()
	if claims.VerifyExpiresAt(now, true) && claims.VerifyIssuer(*url, true) {
		return nil
	}
	return fmt.Errorf("token is invalid")
}

func (claims JwtClaims) VerifyAudience(origin string) bool {
	return strings.Compare(claims.Audience, origin) == 0
}