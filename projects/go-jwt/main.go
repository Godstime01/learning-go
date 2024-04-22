package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var secretKey = []byte("secret-key")

var DB *gorm.DB

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	// Hash the password before saving
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	return nil
}

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func loginHandler(c *gin.Context) {

	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	//
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	//Find user in database
	var user User
	if err := DB.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Compare the hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// create jwt token
	token, err := createToken(user.Username)
	if err != nil {
		// If an error occurs, return a JSON response with status 400 and the error message
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Authentication successful
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

func registerHandler(c *gin.Context) {
	var user User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, user)
}

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	// making migrations
	DB.AutoMigrate(&User{})
}

func main() {

	r := gin.Default()

	// define router
	r.GET("/", func(c *gin.Context) {

		// Retrieve the JWT token from the request headers
		tokenString := c.Request.Header.Get("Authorization")

		// Check if the token is empty or doesn't start with "Bearer "
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			// If the token is missing or invalid, return a JSON response with status 401
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		// Extract the token string without the "Bearer " prefix
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		err := verifyToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to our application",
		})
	})

	r.POST("/register", registerHandler)

	r.POST("/login", loginHandler)
	r.Run()
}
