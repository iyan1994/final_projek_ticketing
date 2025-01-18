package service

import (
	"database/sql"
	"final_projek_ticketing/model"
	"final_projek_ticketing/repository"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateAssignTicketHandler(c *gin.Context) {
	var assignTicketDto model.AssignTicketDto
	username, exists := c.Get("username") // mengambil username dari context
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username not found in context"})
		return
	}
	err := c.ShouldBind(&assignTicketDto)
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

	assignTicket := assignTicketDto.ToModel()

	assignTicket.IdAdmin = user.IdUser
	err = repository.Db.Create(&assignTicket).Error

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to save aasign ticket : %s", err.Error())),
		)
		return
	}
	assignTicketDto.IdTicket = assignTicket.IdTicket
	assignTicketDto.CreatedAt = assignTicket.CreatedAt
	assignTicketDto.UpdatedAt = assignTicket.UpdatedAt

	assignTicketDto.FillFromModel(assignTicket)

	c.JSON(http.StatusOK, model.NewSuccessResponse("success to save assign ticket", assignTicketDto))
}

func StartTicketHandler(c *gin.Context) {
	var assignTicketDto model.AssignTicketDto
	err := c.ShouldBind(&assignTicketDto)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("failed to bind request: %s", err.Error())),
		)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))

	assignTicket := assignTicketDto.ToModel()
	// cek assign ticket by id
	err = repository.Db.First(&assignTicket, id).Error
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			model.NewFailedResponse(fmt.Sprintf("not exist id = %s", err.Error())),
		)
		return
	}

	// Simpan start ticket  ke table assign_ticket
	startTicket := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	assignTicket.StartTicket = startTicket

	err = repository.Db.Save(&assignTicket).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to update start time: %s", err.Error())),
		)
		return
	}

	var ticket model.Ticket
	// cek ticket by id
	err = repository.Db.First(&ticket, assignTicket.IdTicket).Error
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			model.NewFailedResponse(fmt.Sprintf("not exist id = %s", err.Error())),
		)
		return
	}

	// ubah status menjadi "On Going" di table ticket

	ticket.Status = "On Going"

	err = repository.Db.Save(&ticket).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to update status on going ticket: %s", err.Error())),
		)
		return
	}

	assignTicketDto.IdAssignTicket = assignTicket.IdAssignTicket
	assignTicketDto.CreatedAt = assignTicket.CreatedAt
	assignTicketDto.UpdatedAt = assignTicket.UpdatedAt
	assignTicketDto.StartTicket = &startTicket.Time

	assignTicketDto.FillFromModel(assignTicket)

	c.JSON(http.StatusOK, model.NewSuccessResponse("success start ticket", assignTicketDto))

}

func ClosedTicketHandler(c *gin.Context) {
	var assignTicketDto model.AssignTicketDto
	err := c.ShouldBind(&assignTicketDto)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("failed to bind request: %s", err.Error())),
		)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))

	assignTicket := assignTicketDto.ToModel()
	// cek assign ticket by id
	err = repository.Db.First(&assignTicket, id).Error
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			model.NewFailedResponse(fmt.Sprintf("not exist id = %s", err.Error())),
		)
		return
	}

	// Simpan closed ticket  ke table assign_ticket
	closeTicket := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	// Menghitung selisih waktu
	duration := assignTicket.StartTicket.Time.Sub(closeTicket.Time)
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60

	selisihWaktu := fmt.Sprintf("%d hours - %d minutes", hours, minutes)
	finishTime := sql.NullString{
		String: selisihWaktu,
		Valid:  true,
	}

	assignTicket.CloseTicket = closeTicket
	assignTicket.FinishTime = finishTime

	err = repository.Db.Save(&assignTicket).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to update close time: %s", err.Error())),
		)
		return
	}

	var ticket model.Ticket
	// cek ticket by id
	err = repository.Db.First(&ticket, assignTicket.IdTicket).Error
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			model.NewFailedResponse(fmt.Sprintf("not exist id = %s", err.Error())),
		)
		return
	}

	// ubah status menjadi "Closed" di table ticket

	ticket.Status = "Closed"

	err = repository.Db.Save(&ticket).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to update status closed ticket: %s", err.Error())),
		)
		return
	}

	assignTicketDto.IdAssignTicket = assignTicket.IdAssignTicket
	assignTicketDto.CreatedAt = assignTicket.CreatedAt
	assignTicketDto.UpdatedAt = assignTicket.UpdatedAt
	assignTicketDto.StartTicket = &closeTicket.Time
	assignTicketDto.FinishTime = &finishTime.String

	assignTicketDto.FillFromModel(assignTicket)

	c.JSON(http.StatusOK, model.NewSuccessResponse("success closed ticket", assignTicketDto))

}

// view assignticket engineer
func ViewAssignTicketEngineerHandler(c *gin.Context) {
	var assignTicket []model.AssignTicket
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

	query := repository.Db.Model(&model.AssignTicket{})

	// Filter berdasarkan ID teknisi
	query = query.Where("id_teknisi = ?", user.IdUser)

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

	err = query.Limit(req.PageSize).Offset(offset).Find(&assignTicket).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("error select: %s", err.Error())),
		)
		return
	}

	var assignticketsDto []model.AssignTicketDto
	for _, data := range assignTicket {
		var assignticketDto model.AssignTicketDto
		assignticketDto.FillFromModel(data)
		assignticketsDto = append(assignticketsDto, assignticketDto)
	}

	//buat interface untuk meta data
	metaData := map[string]interface{}{
		"page":      req.Page,
		"page_size": req.PageSize,
		"pages":     (total + int64(req.PageSize) - 1) / int64(req.PageSize),
	}

	c.JSON(http.StatusOK, model.NewSuccessResponseView("success", assignticketsDto, metaData))

}

func ViewAssignTicketEngineerByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("invalid id : %s", err.Error())),
		)
		return
	}

	var assignTicket model.AssignTicket

	err = repository.Db.First(&assignTicket, id).Error

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			model.NewFailedResponse(fmt.Sprintf("failed to get : %s", err.Error())),
		)
		return
	}

	var assignTicketDto model.AssignTicketDto
	assignTicketDto.FillFromModel(assignTicket)

	c.JSON(
		http.StatusOK,
		model.NewSuccessResponse("success", assignTicketDto),
	)

}
