package service

import (
	"final_projek_ticketing/controller"
	"final_projek_ticketing/model"
	"final_projek_ticketing/repository"
	"fmt"
	"net/http"
	"strconv"
	"time"

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

// view myticket client
func ViewMyTicketHandler(c *gin.Context) {
	var ticket []model.Ticket
	var total int64

	username, exists := c.Get("username") // mengambil username dari context
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username not found in context"})
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

	var req model.Request
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("error: %s", err.Error())),
		)
		return
	}

	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 {
		req.PageSize = 10
	}

	offset := (req.Page - 1) * req.PageSize

	query := repository.Db.Model(&model.Ticket{})

	// Filter berdasarkan ID
	query = query.Where("id_user = ?", user.IdUser)

	// Filter berdasarkan status
	if req.Status != "" {
		query = query.Where("status = ?", req.Status)
	}

	// Filter berdasarkan rentang tanggal
	if req.Start != "" && req.End != "" {
		start, err1 := time.Parse("2006-01-02", req.Start)
		end, err2 := time.Parse("2006-01-02", req.End)
		if err1 == nil && err2 == nil {
			query = query.Where("created_at BETWEEN ? AND ?", start, end)
		}
	}

	// Hitung total data dan ambil data dengan pagination
	err := query.Count(&total).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("error: %s", err.Error())),
		)
		return
	}

	err = query.Limit(req.PageSize).Offset(offset).Find(&ticket).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("error select: %s", err.Error())),
		)
		return
	}

	var ticketsDto []model.TicketDto
	for _, data := range ticket {
		var ticketDto model.TicketDto
		ticketDto.FillFromModel(data)
		ticketsDto = append(ticketsDto, ticketDto)
	}

	//buat interface untuk meta data
	metaData := map[string]interface{}{
		"page":      req.Page,
		"page_size": req.PageSize,
		"pages":     (total + int64(req.PageSize) - 1) / int64(req.PageSize),
	}

	c.JSON(http.StatusOK, model.NewSuccessResponseView("success", ticketsDto, metaData))

}

// view all ticket untuk admin
func ViewAllTicketHandler(c *gin.Context) {
	var ticket []model.Ticket
	var total int64

	username, exists := c.Get("username") // mengambil username dari context
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username not found in context"})
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

	var req model.Request
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("error: %s", err.Error())),
		)
		return
	}

	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 {
		req.PageSize = 10
	}

	offset := (req.Page - 1) * req.PageSize

	query := repository.Db.Model(&model.Ticket{})

	// Filter berdasarkan status
	if req.Status != "" {
		query = query.Where("status = ?", req.Status)
	}

	// Filter berdasarkan rentang tanggal
	if req.Start != "" && req.End != "" {
		start, err1 := time.Parse("2006-01-02", req.Start)
		end, err2 := time.Parse("2006-01-02", req.End)
		if err1 == nil && err2 == nil {
			query = query.Where("created_at BETWEEN ? AND ?", start, end)
		}
	}

	// Hitung total data dan ambil data dengan pagination
	err := query.Count(&total).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("error: %s", err.Error())),
		)
		return
	}

	err = query.Limit(req.PageSize).Offset(offset).Find(&ticket).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("error select: %s", err.Error())),
		)
		return
	}

	var ticketsDto []model.TicketDto
	for _, data := range ticket {
		var ticketDto model.TicketDto
		ticketDto.FillFromModel(data)
		ticketsDto = append(ticketsDto, ticketDto)
	}

	//buat interface untuk meta data
	metaData := map[string]interface{}{
		"page":      req.Page,
		"page_size": req.PageSize,
		"pages":     (total + int64(req.PageSize) - 1) / int64(req.PageSize),
	}

	c.JSON(http.StatusOK, model.NewSuccessResponseView("success", ticketsDto, metaData))

}

func ViewTicketByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("invalid id : %s", err.Error())),
		)
		return
	}

	var ticket model.Ticket

	err = repository.Db.First(&ticket, id).Error

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			model.NewFailedResponse(fmt.Sprintf("failed to get : %s", err.Error())),
		)
		return
	}

	var ticketDto model.TicketDto
	ticketDto.FillFromModel(ticket)

	c.JSON(
		http.StatusOK,
		model.NewSuccessResponse("success", ticketDto),
	)

}
