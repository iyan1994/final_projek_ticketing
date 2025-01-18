package service

import (
	"database/sql"
	"final_projek_ticketing/controller"
	"final_projek_ticketing/model"
	"final_projek_ticketing/repository"
	"fmt"
	"net/http"
	"strconv"

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

// get client
func GetClientHandler(c *gin.Context) {
	var user []model.User

	err := repository.Db.Where("id_role = 2").Find(&user).Error

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("Failed to get user client: %s", err.Error())),
		)
		return
	}

	var usersDto []model.UserDto
	for _, data := range user {
		var userDto model.UserDto
		userDto.FillFromModel(data)
		usersDto = append(usersDto, userDto)
	}

	c.JSON(http.StatusOK, model.NewSuccessResponse("success", usersDto))
}

// get Engineer
func GetEngineerHandler(c *gin.Context) {
	var user []model.User

	err := repository.Db.Where("id_role = 3").Find(&user).Error

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("Failed to get user engineer: %s", err.Error())),
		)
		return
	}

	var usersDto []model.UserDto
	for _, data := range user {
		var userDto model.UserDto
		userDto.FillFromModel(data)
		usersDto = append(usersDto, userDto)
	}

	c.JSON(http.StatusOK, model.NewSuccessResponse("success", usersDto))
}

func DeleteUserHandler(c *gin.Context) {
	var existUser model.User
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("invalid id : %s", err.Error())),
		)
		return
	}

	// cek ticket apa id user sudah buat
	var ticket []model.Ticket
	err = repository.Db.Where("id_user = ?", id).Find(&ticket).Error
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			model.NewFailedResponse(fmt.Sprintf("not exist id : %s", err.Error())),
		)
		return
	}

	if len(ticket) > 0 {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("id is in another table")),
		)
		return
	}

	// cek assign ticket  sudah buat oleh id user yg di hapus
	var assignTicket []model.AssignTicket
	err = repository.Db.Where("id_teknisi = ?", id).Find(&assignTicket).Error
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			model.NewFailedResponse(fmt.Sprintf("not exist id : %s", err.Error())),
		)
		return
	}

	if len(assignTicket) > 0 {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("id is in another table")),
		)
		return
	}

	//jika id tidak ada, valid untuk delete user
	err = repository.Db.First(&existUser, id).Error
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			model.NewFailedResponse(fmt.Sprintf("not exist id : %s", err.Error())),
		)
		return
	}
	user := model.User{IdUser: id}
	result := repository.Db.Delete(user)

	if result.Error != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to delete user : %s", err.Error())),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		model.NewSuccessResponse(fmt.Sprintf("%d user delete", result.RowsAffected), nil),
	)

}
