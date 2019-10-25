package ginplus

import (
	"fmt"
	"net/http"

	. "github.com/freelifer/litelib/public"
	"github.com/gin-gonic/gin"

	// "encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	GINPLUS_X_REQUEST_ID = "XRequestId"
	SecretKey            = "litelib ---------"
)

// 一些常量
var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := RandomString(6)
		c.Header("X-Request-Id", id)
		c.Next()
	}
}

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(403, gin.H{
				"message": "token is not found",
			})
			c.Abort()
			return
		}

		// parseToken 解析token包含的信息
		claims, err := ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}

type CustomClaims struct {
	jwt.StandardClaims
}

func CreateToken() (string, error) {
	// token := jwt.New(jwt.SigningMethodHS256)
	// claims := make(jwt.MapClaims)
	// claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	// claims["iat"] = time.Now().Unix()
	// token.Claims = claims
	claims := CustomClaims{
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SecretKey))
}

// 解析Tokne
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

func Route(group gin.IRoutes, path string, ci ControllerInterface) {
	group.GET(fmt.Sprintf("/%s", path), ci.List)
	group.GET(fmt.Sprintf("/%s/:id", path), ci.Get)
	group.POST(fmt.Sprintf("/%s", path), ci.Post)
	group.PUT(fmt.Sprintf("/%s/:id", path), ci.Put)
}

func Route2(group gin.IRoutes, path string, id string, ci ControllerInterface) {
	group.GET(fmt.Sprintf("/%s", path), ci.List)
	group.GET(fmt.Sprintf("/%s/:%s", path, id), ci.Get)
	group.POST(fmt.Sprintf("/%s", path), ci.Post)
	group.PUT(fmt.Sprintf("/%s/:%s", path, id), ci.Put)
}
