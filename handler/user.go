package handler

import (
	"crowdfund/helper"
	"crowdfund/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService users.Service
}

func NewUserHandler(userService users.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input users.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	user, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	formatter := users.FormatUser(user, "tokentoken")
	response := helper.APIResponse("Account has been Registered", http.StatusOK, "succes", formatter)

	c.JSON(http.StatusOK, response)
}
