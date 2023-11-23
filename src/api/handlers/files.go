package handlers

import (
	"fmt"
	"github.com/atefeh-syf/car-sale/api/dto"
	_ "github.com/atefeh-syf/car-sale/api/dto"
	"github.com/atefeh-syf/car-sale/api/helper"
	_ "github.com/atefeh-syf/car-sale/api/helper"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/pkg/logging"
	"github.com/atefeh-syf/car-sale/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type FileHandler struct {
	service *services.FileService
}

func NewFileHandler(cfg *config.Config) *FileHandler {
	return &FileHandler{
		service: services.NewFileService(cfg),
	}
}

func (h *FileHandler) Create(c *gin.Context) {
	upload := dto.UploadFileRequest{}
	err := c.ShouldBind(&upload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	req := dto.CreateFileRequest{}
	req.Description = upload.Description
	req.MimeType = upload.File.Header.Get("Content-Type")
	req.Directory = "uploads"
	req.Name, err = saveUploadFile(upload.File, req.Directory)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	res, err := h.service.Create(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, helper.Success))

}

func (h *FileHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

func (h *FileHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, helper.ValidationError))
		return
	}
	file, err := h.service.GetById(c, id)
	if err != nil {
		logger.Error(logging.Io, logging.RemoveFile, err.Error(), nil)
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, helper.NotFoundError))
		return
	}
	err = os.Remove(fmt.Sprintf("%s/%s", file.Directory, file.Name))
	if err != nil {
		logger.Error(logging.Io, logging.RemoveFile, err.Error(), nil)
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, helper.InternalError))
		return
	}
	err = h.service.Delete(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(nil, true, helper.Success))
}

func (h *FileHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

func (h *FileHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFiler)
}

func saveUploadFile(file *multipart.FileHeader, directory string) (string, error) {
	// test.txt -> 95239855629856.txt
	randFileName := uuid.New()
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return "", err
	}
	fileName := file.Filename
	fileNameArr := strings.Split(fileName, ".")
	fileExt := fileNameArr[len(fileNameArr)-1]
	fileName = fmt.Sprintf("%s.%s", randFileName, fileExt)
	dst := fmt.Sprintf("%s/%s", directory, fileName)

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return "", err
	}
	return fileName, nil
}
