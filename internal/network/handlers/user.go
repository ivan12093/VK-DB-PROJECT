package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan12093/VK-DB-PROJECT/internal/models"
	"github.com/ivan12093/VK-DB-PROJECT/internal/usecase"
	"github.com/ivan12093/VK-DB-PROJECT/internal/utils/errors"
	"github.com/mailru/easyjson"
)

type HandlerUsers struct {
	UseCase usecase.UserU
}

func NewUsersHandler(useCase usecase.UserU) *HandlerUsers {
	return &HandlerUsers{UseCase: useCase}
}

func (handler *HandlerUsers) Get(c *gin.Context) {
	nickname := c.Param("nickname")
	model, err := handler.UseCase.Get(&nickname)
	if err != nil {
		c.AbortWithStatusJSON(err.(errors.MsgErrors).Code(), err)
		return
	}
	c.JSON(http.StatusOK, model)
}

func (handler *HandlerUsers) Create(c *gin.Context) {
	model := &models.User{}
	model.Username = c.Param("nickname")

	err := easyjson.UnmarshalFromReader(c.Request.Body, model)
	if err != nil {
		c.AbortWithStatusJSON(errors.BadRequest.Code(), errors.BadRequest)
		return
	}

	users, err := handler.UseCase.Create(model)
	if users != nil {
		c.JSON(err.(errors.MsgErrors).Code(), users)
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(err.(errors.MsgErrors).Code(), err)
		return
	}

	c.JSON(http.StatusCreated, model)
}

func (handler *HandlerUsers) Update(c *gin.Context) {
	model := &models.User{}
	model.Username = c.Param("nickname")
	err := easyjson.UnmarshalFromReader(c.Request.Body, model)
	if err != nil {
		c.AbortWithStatusJSON(errors.BadRequest.Code(), errors.BadRequest)
		return
	}

	user, err := handler.UseCase.Update(model)

	if err != nil {
		c.AbortWithStatusJSON(err.(errors.MsgErrors).Code(), err)
		return
	}

	c.JSON(http.StatusOK, user)
}
