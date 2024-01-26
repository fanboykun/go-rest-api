package app
	import (
		jwt "github.com/dgrijalva/jwt-go"
	)

type Token struct {
	Success bool `json:"status"`
	Token string `json:"token"`
}

func Generate() Token {
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		return Token{
			Success: false,
			Token: "",
		}
	}
	return Token{
		Success: true,
		Token: token,
	}
}