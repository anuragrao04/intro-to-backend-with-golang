package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var SECRET_KEY = "I like icecream"

var users = []User{
	{
		Username: "anuragrao04",
		Password: "supersecurepassword",
		Role:     "admin",
	},
	{
		Username: "ilikecupcakes",
		Password: "cupcakes123",
		Role:     "user",
	},
}

func main() {
	router := gin.Default()
	router.POST("/login", login)
	router.POST("/verify", verify)
	router.Run(":6969")
}

type loginRequestFormat struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// this function issues a JWT with a TTL of 30 days
func login(c *gin.Context) {
	var req loginRequestFormat
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": "bad request",
		})
		return
	}

	for _, user := range users {
		if user.Username == req.Username && user.Password == req.Password {
			jwtString, err := getJWT(user)
			if err != nil {
				c.JSON(500, gin.H{
					"error": "internal server error",
				})
			} else {
				c.JSON(200, gin.H{
					"token": jwtString,
				})
			}
			return
		}
	}

	c.JSON(401, gin.H{
		"error": "invalid credentials",
	})

}

func getJWT(user User) (string, error) {
	ttl := time.Hour * 24 * 30
	claims := jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(ttl).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Println("error signing token", err)
		return "", err
	}
	return tokenString, nil
}

func verify(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(401, gin.H{
			"error": "missing token",
		})
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		c.JSON(401, gin.H{
			"error": "invalid token",
		})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.JSON(200, gin.H{
			"username": claims["username"],
			"role":     claims["role"],
		})
	} else {
		c.JSON(401, gin.H{
			"error": "invalid token",
		})
	}
}
