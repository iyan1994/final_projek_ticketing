package service

import (
	"final_projek_ticketing/model"
	"final_projek_ticketing/repository"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var imageTicketUploadDir = "uploads/imageticket"

func UploadImagesTicketHandler(c *gin.Context) {
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

	// Ambil file dari form
	formFile, file, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("failed to get image : %s", err.Error())),
		)
		return
	}

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("failed to get upload file : %s", err.Error())),
		)
		return
	}
	defer formFile.Close()

	maxFileSize := 5 * 1024 * 1024 // 5 MB
	if file.Size > int64(maxFileSize) {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse("file size exceed maximum"),
		)
		return

	}

	ext := filepath.Ext(file.Filename)
	if ext != ".jpg" && ext != ".jpeg" {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse("please upload jpg or jpeg"),
		)
		return
	}

	buffer := make([]byte, 512)
	_, err = formFile.Read(buffer)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse("failed to read file buffer"),
		)
		return
	}

	log.Println("Checking file mime type")

	// Detect the MIME type
	mimeType := http.DetectContentType(buffer)
	if !strings.Contains(mimeType, "image/jpeg") {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse("file is not a jpeg picture"),
		)
		return
	}

	// Ambil name image dari form
	name_image := c.PostForm("name_image")
	// Ambil deskripsi dari form
	description := c.PostForm("description")
	path := filepath.Join(imageTicketUploadDir, name_image)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to save file: %s", err.Error())),
		)
		return
	}

	var ticketimage model.TicketImage

	ticketimage.IdTicket = id
	ticketimage.Image = path
	ticketimage.Deksripsi = description
	err = repository.Db.Create(&ticketimage).Error

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to save image ticket : %s", err.Error())),
		)
		return
	}

	c.JSON(http.StatusOK, model.NewSuccessResponse("success to save image ticket", ticketimage))
}

func DownloadImageTicket(c *gin.Context) {
	filename := c.Param("filename")
	imagePath := filepath.Join("uploads", "imageticket", filename)
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		c.JSON(
			http.StatusNotFound,
			model.NewFailedResponse(fmt.Sprintf("image not found: %s", err.Error())),
		)
		return
	}

	// Serve the image file
	c.File(imagePath)
}

// view image ticket
func ViewImageTicketByIdTicket(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("invalid id : %s", err.Error())),
		)
		return
	}

	var ticketImage model.TicketImage

	err = repository.Db.Where("id_ticket = ?", id).Find(&ticketImage).Error

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			model.NewFailedResponse(fmt.Sprintf("failed to get : %s", err.Error())),
		)
		return
	}

	var ticketImageDto model.TicketImageDto
	ticketImageDto.FillFromModel(ticketImage)

	c.JSON(
		http.StatusOK,
		model.NewSuccessResponse("success", ticketImageDto),
	)
}
