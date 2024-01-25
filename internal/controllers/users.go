package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josh1248/forum-website-backend/internal/auth"
	"github.com/josh1248/forum-website-backend/internal/entities"
	"github.com/josh1248/forum-website-backend/internal/models"
)

func GetAllUsers(c *gin.Context) {
	users, err := models.FindAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func GetUserByName(c *gin.Context) {
	user, err := models.FindUserByName(c.Param("name"))
	log.Println(err)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, err)
	} else if err != nil {
		//placeholder error code here, for anything else not successful.
		//consider cases: server error?
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, user)
	}

}

// check if input data is correct. if true, return a cookie.
func VerifyUser(c *gin.Context) {
	var loginUser entities.InputUser

	err := c.ShouldBindJSON(&loginUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, "faulty post data.")
		return
	}

	ok, err := models.AuthenticateUser(loginUser)
	if err == sql.ErrNoRows || !ok {
		c.JSON(http.StatusUnauthorized, "Wrong username or password.")
		return
	} else if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}

	jwt_token, err := auth.GenerateJWT(loginUser.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	//3rd argument sets time of validity of cookie.
	//negative number to indicate transient cookie.
	//4th argument indicates where cookie is valid, 5th argument for routes.
	//use blank for valid cookie throughout any place.
	//5th argument sets toggle to allow only HTTPS. false for testing purposes.
	c.SetCookie("jwt", jwt_token, 60*60*24, "/", "", false, false)
}

func CheckCookie(c *gin.Context) {

	valid, err := auth.VerifyJWT(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{"message": "something went wrong."})
	} else if !valid {
		c.JSON(http.StatusUnauthorized, map[string]any{"message": "Invalid JWT."})
	} else {
		c.JSON(http.StatusOK, map[string]any{"message": "JWT verified."})
	}
}

// weigh server-side vs client-side data validation. security vs ease of implementation.
func CreateUser(c *gin.Context) {
	var newUser entities.InputUser

	//Unmarshals JSON data and reads into struct
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}

	err = models.CreateUser(newUser)
	if err != nil {
		//placeholder controls.
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, "successfully created.")
		return
	}
}
