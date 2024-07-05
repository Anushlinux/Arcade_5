package controllers

import (
	"gopro/initializers"
	"gopro/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// Implement the sign up logic here
	//get email and password of req body

	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	//hash the pass1word
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to hash password",
		})
	}

	//create a new user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create user",
		})
		return
	}
	//respond
	c.JSON(http.StatusCreated, gin.H{})

}

// func SignIn(c *gin.Context) {
// 	// Implement the sign in logic here
// 	//get email and password of req body

// 	var body struct {
// 		Email    string
// 		Password string
// 	}

// 	if c.Bind(&body) != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "bad request",
// 		})
// 		return
// 	}

// 	//get the user from the db
// 	var user models.User
// 	initializers.DB.First("email = ?", body.Email)

// 	if user.ID == 0 {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": user.ID,
// 		})
// 		return
// 	}

// 	//compare the password
// 	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": "invalid password",
// 		})
// 		return
// 	}

// 	//generate a jwt token
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"sub": user.ID,
// 		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
// 	})
// 	tokenString, err := token.SignedString([]byte(os.Getenv("JWT")))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "failed to generate token",
// 		})
// 		return
// 	}

// 	//respond
// 	c.JSON(http.StatusOK, gin.H{
// 		"token": tokenString,
// 	})

// }

// login with jwt token
func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}

	//query
	var user models.User
	initializers.DB.Where("email =?", body.Email).First(&user)

	if user.ID == 0 {
		// Check if the email address is correctly stored in the database
		// Check if the user account is active
		// Check if there are any typos or formatting issues with the email address
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid email",
		})
		return
	}

	//compare the password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid password",
		})
		return
	}

	//token generation
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed token generation failed",
		})
		return
	}

	//set jwt token as cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", tokenString, 3600*24*30, "", "", false, false)
	c.JSON(http.StatusOK, gin.H{})

}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}