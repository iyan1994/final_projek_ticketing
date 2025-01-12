package service

import (
	"final_projek_ticketing/controller"
	"final_projek_ticketing/model"
	"final_projek_ticketing/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTicketHandler(c *gin.Context) {
	var ticketDto model.TicketDto
	username, exists := c.Get("username") // mengambil username dari context
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username not found in context"})
		return
	}
	err := c.ShouldBind(&ticketDto)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("failed to bind request: %s", err.Error())),
		)
		return
	}

	var user model.User
	// cek user by username
	if err := repository.Db.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(
			http.StatusUnauthorized,
			model.NewFailedResponse(fmt.Sprintf("Invalid user: %s", err.Error())),
		)
		return
	}

	ticket := ticketDto.ToModel()
	ticketNumber, err := controller.GenerateTicketNumber(repository.Db)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("Failed to generate ticket number: %s", err.Error())),
		)
		return
	}

	ticket.NoTiket = string(ticketNumber)
	ticket.Status = "Open"
	ticket.IdUser = user.IdUser
	err = repository.Db.Create(&ticket).Error

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to save ticket : %s", err.Error())),
		)
		return
	}
	ticketDto.IdTicket = ticket.IdTicket
	ticketDto.CreatedAt = ticket.CreatedAt
	ticketDto.UpdatedAt = ticket.UpdatedAt

	ticketDto.FillFromModel(ticket)

	c.JSON(http.StatusOK, model.NewSuccessResponse("success to save ticket", ticketDto))
}
