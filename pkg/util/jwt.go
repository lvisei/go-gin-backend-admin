package util

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	// errMissingHeader means the `Authorization` header was empty.
	errMissingHeader = errors.New("The length of the `Authorization` header is zero.")
	jwtSecret        string
)

// TokenContext is the Tokencontext of the JSON web token.
type TokenContext struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	TokenContext
	jwt.StandardClaims
}

// GenerateToken signs the Tokencontext with the specified secret.
func GenerateToken(ctx TokenContext, secret string) (token string, err error) {
	if secret == "" {
		secret = jwtSecret
	}

	nowTime := time.Now()
	expireTime := nowTime.Add(1 * time.Hour / 2)

	claims := Claims{
		TokenContext{
			ctx.ID,
			ctx.Username,
			ctx.Password,
		},
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "go-gin-backend-admin",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = tokenClaims.SignedString([]byte(secret))

	return token, err
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}
}

// ParseToken validates the token with the specified secret,
// and returns the Tokencontext if the token was valid.
func ParseToken(token string, secret string) (*TokenContext, error) {
	ctx := &TokenContext{}

	if secret == "" {
		secret = jwtSecret
	}

	// Parse the token.
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, secretFunc(secret))

	// Parse error.
	if err != nil {
		return ctx, err

		// Read the token if it's valid.
	} else if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		ctx.ID = claims.TokenContext.ID
		ctx.Username = claims.TokenContext.Username
		ctx.Password = claims.TokenContext.Password
		return ctx, nil

		// Other errors.
	} else {
		return ctx, err
	}
}

// ParseRequestToken gets the token from the header and
// pass it to the Parse function to parses the token.
func ParseRequestToken(c *gin.Context) (*TokenContext, error) {
	token := c.Request.Header.Get("Authorization")

	if len(token) == 0 {
		return &TokenContext{}, errMissingHeader
	}

	return ParseToken(token, jwtSecret)
}
