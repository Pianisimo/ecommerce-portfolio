package controllers

import (
	"backend/database"
	"backend/models"
	"backend/myJwt"
	"backend/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

var (
	Validate = validator.New()
)

// SignUp request should receive first_name, last_name, email and password
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelFunc()

		var user models.User
		err := c.BindJSON(&user)
		if err != nil {
			fmt.Println("error binding request to user model")
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		err = Validate.Struct(user)
		if err != nil {
			fmt.Println("error validating user model from request")
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		// Check if user email already exists in the database
		_, err = database.GetUserByEmail(user.Email)
		if err == nil {
			fmt.Println("user email already in use, use another email")
			c.String(http.StatusBadRequest, "user email already in use, use another email.")
			return
		}

		hash := utils.HashPassword(user.Password)
		user.Password = hash

		user.ID = database.CreateUser(user)

		authTokenString, refreshTokenString, csrfSecret, err := myJwt.CreateNewTokens(user.ID, "user")
		if err != nil {
			fmt.Println("error creating tokens for the new user")
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Do not send Cookies data to frontend, use httponly cookies instead.
		user.Token = ""
		user.RefreshToken = ""
		user.CSRF = ""
		user.Password = ""

		err = database.StoreUserTokens(user.ID, authTokenString, refreshTokenString, csrfSecret)
		if err != nil {
			fmt.Println("could not create tokens for the user")
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.SetCookie("token", authTokenString, 3600, "/", "localhost", true, true)
		c.SetCookie("refresh_token", refreshTokenString, 3600, "/", "localhost", true,
			true)
		c.SetCookie("csrf", csrfSecret, 3600, "/", "localhost", true, true)
		c.JSON(http.StatusCreated, user)
	}
}

// Login request should receive email and password
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelFunc()

		var user models.User
		err := c.BindJSON(&user)
		if err != nil {
			fmt.Println("error binding request to user model")
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		// Check if user email already exists in the database
		storedUser, err := database.GetUserByEmail(user.Email)
		if err != nil {
			fmt.Println("email not found in database")
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		// Compare password with stored password(hashed)
		_, err = utils.IsPasswordValid(storedUser.Password, user.Password)
		if err != nil {
			fmt.Println(err.Error())
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		// Refresh user tokens
		authTokenString, refreshTokenString, csrfSecret, err := myJwt.CreateNewTokens(storedUser.ID, "user")
		if err != nil {
			fmt.Println("error creating new tokens for the user")
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Do not send Cookies data to frontend, use httponly cookies instead.
		storedUser.Token = ""
		storedUser.RefreshToken = ""
		storedUser.CSRF = ""
		storedUser.Password = ""

		err = database.StoreUserTokens(storedUser.ID, authTokenString, refreshTokenString, csrfSecret)
		if err != nil {
			fmt.Println("could not create tokens for the user")
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.SetCookie("token", authTokenString, 3600, "/", "localhost", true, true)
		c.SetCookie("refresh_token", refreshTokenString, 3600, "/", "localhost", true, true)
		c.SetCookie("csrf", csrfSecret, 3600, "/", "localhost", true, true)
		c.JSON(http.StatusOK, storedUser)
	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelFunc()

		tokenString, err := c.Cookie("token")
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		refreshTokenString, err := c.Cookie("refresh_token")
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		token, refreshToken, err := myJwt.RevokeTokens(tokenString, refreshTokenString)
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.SetCookie("token", token, 0, "/", "localhost", true, true)
		c.SetCookie("refresh_token", refreshToken, 0, "/", "localhost", true,
			true)
		c.SetCookie("csrf", "", 0, "/", "localhost", true, true)
		c.String(http.StatusOK, "successfully logged out")
	}
}

// IsAuth checks if client has valid cookies and returns a user from cookies values.
func IsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelFunc()

		authTokenString, err := c.Cookie("token")
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		refreshTokenString, err := c.Cookie("refresh_token")
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		csrf, err := c.Cookie("csrf")
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		valid := myJwt.CheckIfValid(authTokenString, refreshTokenString, csrf)

		if valid {
			user, err := myJwt.GetUserFromAuthTokenString(authTokenString)
			if err != nil {
				fmt.Println(err)
				c.String(http.StatusInternalServerError, err.Error())
				return
			}

			c.JSON(http.StatusOK, user)
			return
		} else {
			c.String(http.StatusUnauthorized, "unauthorized")
			return
		}
	}
}

/*
func ProductViewerAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
		var products models.Product
		defer cancelFunc()
		err := c.BindJSON(&products)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		products.ID = primitive.NewObjectID()
		_, err = ProductCollection.InsertOne(ctx, products)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not inserted"})
			return
		}

		c.JSON(http.StatusOK, "Successfully added")
	}
}
func SearchProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var productList []models.Product
		ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelFunc()

		cursor, err := ProductCollection.Find(ctx, bson.D{{}})
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "something went wrong")
			return
		}

		err = cursor.All(ctx, &productList)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		defer cursor.Close(ctx)
		err = cursor.Err()
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusBadRequest, "invalid")
			return
		}
		defer cancelFunc()

		c.IndentedJSON(http.StatusOK, productList)
	}
}
func SearchProductByQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		var searchedProducts []models.Product
		queryParam := c.Query("name")
		if queryParam == "" {
			log.Println("query is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid search index"})
			return
		}

		ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelFunc()

		searchQueryDb, err := ProductCollection.Find(ctx, bson.M{
			"name": bson.M{"$regex": queryParam},
		})
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, "something went wrong while fetching the data")
			return
		}

		err = searchQueryDb.All(ctx, &searchedProducts)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusBadRequest, "invalid")
			return
		}
		defer searchQueryDb.Close(ctx)
		err = searchQueryDb.Err()
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusBadRequest, "invalid request")
			return
		}
		defer cancelFunc()

		c.IndentedJSON(http.StatusOK, searchedProducts)
	}
}*/
