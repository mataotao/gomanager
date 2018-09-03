package token

import (
	"apiserver/pkg/global/redis"
	"bytes"
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	redisgo "github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

var (
	// ErrMissingHeader means the `Authorization` header was empty.
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")
)

// Context is the context of the JSON web token.
type Context struct {
	ID       uint64
	Username string
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

// Parse validates the token with the specified secret,
// and returns the context if the token was valid.
func Parse(tokenString string, secret string) (*Context, error) {
	ctx := &Context{}

	// Parse the token.
	token, err := jwt.Parse(tokenString, secretFunc(secret))

	// Parse error.
	if err != nil {
		return ctx, err

		// Read the token if it's valid.
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.ID = uint64(claims["id"].(float64))
		ctx.Username = claims["username"].(string)
		uId := strconv.Itoa(int(ctx.ID))
		var key bytes.Buffer
		key.WriteString("user:login:")
		key.WriteString(uId)
		loginStr := key.String()
		pool := redis.Pool.Pool.Get()
		defer pool.Close()

		redisToken, err := redisgo.String(pool.Do("GET", loginStr))
		if err != nil || redisToken != tokenString {
			return ctx, errors.New("key无效")
		}

		var permissionKey bytes.Buffer
		permissionKey.WriteString("user:permission:ids:")
		permissionKey.WriteString(uId)

		var menuKey bytes.Buffer
		menuKey.WriteString("user:menu:")
		menuKey.WriteString(uId)
		//设置有效时间
		if _, err = pool.Do("EXPIRE", loginStr, 20*60); err != nil {
			return nil, err
		}
		if _, err = pool.Do("EXPIRE", permissionKey.String(), 20*60); err != nil {
			return nil, err
		}
		if _, err = pool.Do("EXPIRE", menuKey.String(), 20*60); err != nil {
			return nil, err
		}

		return ctx, nil

		// Other errors.
	} else {
		return ctx, err
	}
}

// ParseRequest gets the token from the header and
// pass it to the Parse function to parses the token.
func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	// Load the jwt secret from config
	secret := viper.GetString("jwt_secret")

	if len(header) == 0 {
		return &Context{}, ErrMissingHeader
	}

	var t string
	// Parse the header to get the token part.
	fmt.Sscanf(header, "Bearer %s", &t)
	return Parse(t, secret)
}

// Sign signs the context with the specified secret.
func Sign(c Context, secret string) (tokenString string, err error) {
	// Load the jwt secret from the Gin config if the secret isn't specified.
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}
	// The token content.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.ID,
		"username": c.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})
	// Sign the token with the specified secret.
	tokenString, err = token.SignedString([]byte(secret))

	return
}
