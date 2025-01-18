package service

import (
	"final_projek_ticketing/model"
	"final_projek_ticketing/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateFeedbackHandler(c *gin.Context) {
	var feedbackDto model.FeedbackDto
	username, exists := c.Get("username") // mengambil username dari context
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username not found in context"})
		return
	}
	err := c.ShouldBind(&feedbackDto)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("failed to bind request: %s", err.Error())),
		)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("invalid id : %s", err.Error())),
		)
		return
	}

	var ticket model.Ticket
	// cek ticket by id
	err = repository.Db.First(&ticket, id).Error
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			model.NewFailedResponse(fmt.Sprintf("not exist id = %s", err.Error())),
		)
		return
	}

	if ticket.Status != "Closed" {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse("Tickets are not yet closed"),
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

	feedback := feedbackDto.ToModel()
	feedback.IdTicket = id
	err = repository.Db.Create(&feedback).Error

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to save feedback : %s", err.Error())),
		)
		return
	}
	feedbackDto.IdFeedback = feedback.IdFeedback
	feedbackDto.CreatedAt = feedback.CreatedAt
	feedbackDto.UpdatedAt = feedback.UpdatedAt
	feedbackDto.IdTicket = feedback.IdTicket

	feedbackDto.FillFromModel(feedback)

	c.JSON(http.StatusOK, model.NewSuccessResponse("success to save feedback", feedbackDto))
}

// view feedback
func ViewFeedbackByIdTicket(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("invalid id : %s", err.Error())),
		)
		return
	}

	var feedback model.Feedback

	err = repository.Db.Where("id_ticket = ?", id).Find(&feedback).Error

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			model.NewFailedResponse(fmt.Sprintf("failed to get : %s", err.Error())),
		)
		return
	}

	var feedbackDto model.FeedbackDto
	feedbackDto.FillFromModel(feedback)

	c.JSON(
		http.StatusOK,
		model.NewSuccessResponse("success", feedbackDto),
	)
}
