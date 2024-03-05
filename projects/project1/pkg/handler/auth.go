package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kirigaikabuto/hero_preparation/projects/project1/models"
	"github.com/kirigaikabuto/hero_preparation/projects/project1/utils"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		utils.NewHttpErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

}

func (h *Handler) logIn(c *gin.Context) {

}
