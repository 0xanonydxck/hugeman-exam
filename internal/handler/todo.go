package handler

import (
	"github.com/dxckboi/hugeman-exam/internal/service"
	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	serv service.TodoService
}

func NewTodoHandler(router *gin.RouterGroup, serv service.TodoService) *todoHandler {
	h := &todoHandler{serv: serv}

	{
		router.GET("", h.All)
		router.GET("/:id", h.Get)
		router.POST("", h.Create)
		router.PUT("/:id", h.Update)
		router.PATCH("/:id/in-progress", h.SetInProgress)
		router.PATCH("/:id/completed", h.SetCompleted)
		router.DELETE("/:id", h.Delete)
	}

	return h
}

// @Summary Get all todo tasks
// @Description retrieves a list of all todo tasks.
// @Tags todo
// @Accept json
// @Produce json
// @Success 200 {object} Response{result=[]service.TodoResponse} "List of todo items"
// @Failure 500 {object} ErrorResponse "Error message"
// @Router /todo [get]
func (h *todoHandler) All(c *gin.Context) {
	result, err := h.serv.All(nil)
	if err != nil {
		ResponseError(c, err)
		return
	}

	ResponseOK(c, result)
}

// @Summary Get todo task by ID
// @Description retrieve todo task by ID.
// @Tags todo
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} Response{result=service.TodoResponse} "List of todo items"
// @Failure 500 {object} ErrorResponse "Error message"
// @Router /todo/{id} [get]
func (h *todoHandler) Get(c *gin.Context) {
	id := c.Param("id")
	result, err := h.serv.Get(id)
	if err != nil {
		ResponseError(c, err)
		return
	}

	ResponseOK(c, result)
}

// @Summary Create todo task
// @Description insert todo task into the database.
// @Tags todo
// @Accept json
// @Produce json
// @Param request body service.CreateTodoRequest true "create todo request body"
// @Success 200 {object} Response{result=service.TodoResponse} "List of todo items"
// @Failure 500 {object} ErrorResponse "Error message"
// @Router /todo [post]
func (h *todoHandler) Create(c *gin.Context) {
	body := service.CreateTodoRequest{}
	if err := c.ShouldBind(&body); err != nil {
		ResponseError(c, err)
		return
	}

	result, err := h.serv.Create(&body)
	if err != nil {
		ResponseError(c, err)
		return
	}

	ResponseOK(c, result)
}

// @Summary Update todo task by ID
// @Description update todo task by ID into the database.
// @Tags todo
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Param request body service.UpdateTodoRequest true "update todo request body"
// @Success 200 {object} Response{result=service.TodoResponse} "List of todo items"
// @Failure 500 {object} ErrorResponse "Error message"
// @Router /todo/{id} [put]
func (h *todoHandler) Update(c *gin.Context) {
	id := c.Param("id")
	body := service.UpdateTodoRequest{}
	if err := c.ShouldBind(&body); err != nil {
		ResponseError(c, err)
		return
	}

	result, err := h.serv.Update(id, &body)
	if err != nil {
		ResponseError(c, err)
		return
	}

	ResponseOK(c, result)
}

// @Summary Set todo task to in-progress
// @Description set selected todo task to in-progress.
// @Tags todo
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} Response "List of todo items"
// @Failure 500 {object} ErrorResponse "Error message"
// @Router /todo/{id}/in-progress [patch]
func (h *todoHandler) SetInProgress(c *gin.Context) {
	id := c.Param("id")
	if err := h.serv.SetInProgress(id); err != nil {
		ResponseError(c, err)
		return
	}

	ResponseOK(c, nil)
}

// @Summary Set todo task to completed
// @Description set selected todo task to completed.
// @Tags todo
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} Response "List of todo items"
// @Failure 500 {object} ErrorResponse "Error message"
// @Router /todo/{id}/completed [patch]
func (h *todoHandler) SetCompleted(c *gin.Context) {
	id := c.Param("id")
	if err := h.serv.SetCompleted(id); err != nil {
		ResponseError(c, err)
		return
	}

	ResponseOK(c, nil)
}

// @Summary Delete todo task by ID
// @Description delete todo task by ID from the database.
// @Tags todo
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} Response "List of todo items"
// @Failure 500 {object} ErrorResponse "Error message"
// @Router /todo/{id} [delete]
func (h *todoHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.serv.Delete(id); err != nil {
		ResponseError(c, err)
		return
	}

	ResponseOK(c, nil)
}
