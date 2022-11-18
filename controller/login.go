package controller

import (
	"net/http"
	"time"

	"webtool-api/middleware"
	"webtool-api/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	loginInput body models.LoginRequest true "Login Input"
// @Success		  	200 {object} models.ResponsetokenModel
// @Failure 		401 {object} models.ResponseWoModel
// @Router			/Login/ [post]
func Login(c *fiber.Ctx) error {
	var dataInput map[string]string

	if parErr := c.BodyParser(&dataInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	var loginModel models.LoginResponse

	middleware.DBConn.Debug().Table("mfs.view_login").Find(&loginModel, dataInput)

	// if loginModel.User_id == 0 {
	// 	return c.Status(http.StatusNotFound).JSON(models.ResponseWoModel{
	// 		RetCode: "401",
	// 		Message: "User not found!",
	// 	})
	// }

	if parErr := bcrypt.CompareHashAndPassword([]byte(loginModel.User_passwd), []byte(dataInput["c.user_passwd"])); parErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "User not found!",
			Data:    parErr.Error(),
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = loginModel.User_login
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(http.StatusCreated).JSON(models.ResponsetokenModel{
		RetCode: "200",
		Message: "Succes",
		Data:    loginModel,
		Token:   t,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Success		  	200 {object} models.ResponseModel
// @Failure 		400 {object} models.ResponsetokenModel
// @Router			/Authenticate/ [post]
func Authenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "unauthenticated",
			Data:    err.Error(),
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var login models.LoginResponse

	middleware.DBConn.Debug().Table("mfs.view_login").Find(&claims.Issuer, login)

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    login,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Success		  	200 {object} models.ResponseWoModel
// @Router			/Logout/ [post]
func Logout(c *fiber.Ctx) error {

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
		RetCode: "200",
		Message: "Succes",
	})
}
