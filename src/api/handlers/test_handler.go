package handlers

import (
	"net/http"

	"github.com/atefeh-syf/car-sale/api/helper"
	"github.com/gin-gonic/gin"
)

type header struct {
	UserId string
	Browser string
}

type personData struct {
	FirstName string
	LastName string
}
type TestHandler struct {

}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"result": "Test",
	})
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("Test", true, 0))

}

func (h *TestHandler) Users(c *gin.Context) {

	c.JSON(http.StatusOK, helper.GenerateBaseResponse("Users", true, 0))

}

func (h *TestHandler) UserById(c *gin.Context) {

	id := c.Param("id")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "UserById",
		"id":     id,
	}, true, 0))

}

func (h *TestHandler) UserByUsername(c *gin.Context) {

	username := c.Param("username")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result":   "UserByUsername",
		"username": username,
	}, true, 0))
}

func (h *TestHandler) Accounts(c *gin.Context) {

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "Accounts",
	}, true, 0))
}

func (h *TestHandler) AddUser(c *gin.Context) {

	id := c.Param("id")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "AddUser",
		"id":     id,
	}, true, 0))
}

func (h *TestHandler) HeaderBinder1 (c *gin.Context) {
	userId := c.GetHeader("UserId")
	c.JSON(http.StatusOK, gin.H{
		"result" : "HeaderBinder1",
		"userId" : userId,
	})
}


func (h *TestHandler) HeaderBinder2 (c *gin.Context) {
	header := header{}
	c.BindHeader(header)
	c.JSON(http.StatusOK, gin.H{
		"result" : "HeaderBinder2",
		"header" : &header,
	})
}

func (h *TestHandler) QueryBinder1 (c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"result" : "QueryBinder1",
		"id" : id,
		"name" : name,
	})
}


func (h *TestHandler) QueryBinder2 (c *gin.Context) {
	ids := c.QueryArray("id") //  ?id=12&name=test&id=19 => ids : [12, 19] , name=test
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"result" : "QueryBinder2",
		"id" : ids,
		"name" : name,
	})
}



func (h *TestHandler) UriBinder (c *gin.Context) {
	ids := c.Param("id") 
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"result" : "UriBinder",
		"id" : ids,
		"name" : name,
	})
}

func (h *TestHandler) BodyBinder(c *gin.Context) {
	p := personData{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil,
				false, -1, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "BodyBinder",
		"person": p,
	}, true, 0))
}


func (h *TestHandler) FormBinder (c *gin.Context) {
	p := personData{}
	c.Bind(&p)
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "FormBinder",
		"person": p,
		}, true, 0))
}


func (h *TestHandler) FileBinder(c *gin.Context) {
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, "file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return	
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "FileBinder",
		"file":   file.Filename,
	}, true, 0))
}