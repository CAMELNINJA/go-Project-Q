package viewmodels

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Username           string `json:"username" :"username"`
	jwt.StandardClaims `:"jwt_._standard_claims"`
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
