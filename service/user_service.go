package service

import (
	"database/sql"
	"final_projek_ticketing/controller"
	"final_projek_ticketing/model"
	"final_projek_ticketing/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserHandler(c *gin.Context) {
	var userDto model.UserDto
	err := c.ShouldBind(&userDto)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("failed to bind request: %s", err.Error())),
		)
		return
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("failed to encrpt password : %s", err.Error())),
		)
		return
	}

	userDto.Password = string(passwordHash)
	user := userDto.ToModel()

	err = repository.Db.Create(&user).Error

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to save user : %s", err.Error())),
		)
		return
	}
	userDto.IdUser = user.IdUser
	userDto.CreatedAt = user.CreatedAt
	userDto.UpdatedAt = user.UpdatedAt

	userDto.FillFromModel(user)

	c.JSON(http.StatusOK, model.NewSuccessResponse("success to save user", userDto))

}

func LoginUserHandler(c *gin.Context) {
	var userDto model.UserDto
	err := c.ShouldBind(&userDto)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("failed to bind request: %s", err.Error())),
		)
		return
	}

	user := userDto.ToModel()
	// cek user by username
	if err := repository.Db.Where("username = ?", userDto.Username).First(&user).Error; err != nil {
		c.JSON(
			http.StatusUnauthorized,
			model.NewFailedResponse(fmt.Sprintf("Invalid username or password: %s", err.Error())),
		)
		return
	}

	// cek user by password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDto.Password)); err != nil {
		c.JSON(
			http.StatusUnauthorized,
			model.NewFailedResponse(fmt.Sprintf("Invalid username or password: %s", err.Error())),
		)
		return
	}

	// Generate token
	token, expireAt, err := controller.GenerateToken(user.Username, user.IdRole, user.Title)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("Failed to generate token: %s", err.Error())),
		)
		return
	}

	// Simpan token ke database
	tokenValue := sql.NullString{
		String: token,
		Valid:  true,
	}
	expiredTokenValue := sql.NullTime{
		Time:  expireAt,
		Valid: true,
	}
	user.Token = tokenValue
	user.ExpiredToken = expiredTokenValue

	err = repository.Db.Save(&user).Error

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to save token: %s", err.Error())),
		)
		return
	}

	userDto.IdUser = user.IdUser
	userDto.CreatedAt = user.CreatedAt
	userDto.UpdatedAt = user.UpdatedAt
	userDto.Token = &token
	userDto.ExpiredToken = &expireAt

	userDto.FillFromModel(user)

	c.JSON(http.StatusOK, model.NewSuccessResponse("success login", userDto))

}
