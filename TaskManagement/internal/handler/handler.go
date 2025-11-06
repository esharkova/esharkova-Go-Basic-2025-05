package handler

import (
	task "TaskManagement/internal/model/task"
	user "TaskManagement/internal/model/user"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc useCase
}

func (h *Handler) DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err = h.uc.DeleteUser(ctx.Request.Context(), id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}

		ctx.Status(http.StatusNoContent)

	}
}

func (h *Handler) UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		args := user.UpdateUserRequest{}

		err = ctx.ShouldBindJSON(&args)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		res, err := h.uc.UpdateUser(ctx.Request.Context(), args, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, res)

	}
}

func (h *Handler) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		args := user.CreateUserRequest{}

		err := ctx.ShouldBindJSON(&args)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		res := h.uc.CreateUser(ctx.Request.Context(), args)
		ctx.JSON(http.StatusCreated, res)

	}
}

func (h *Handler) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		res, err := h.uc.GetUser(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, err.Error())
			return

		}

		ctx.JSON(http.StatusOK, res)

	}
}

func (h *Handler) DeleteTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err = h.uc.DeleteTask(ctx.Request.Context(), id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}

		ctx.Status(http.StatusNoContent)

	}
}

func (h *Handler) UpdateTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		args := task.UpdateTaskRequest{}

		err = ctx.ShouldBindJSON(&args)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		res, err := h.uc.UpdateTask(ctx.Request.Context(), args, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, res)

	}
}

func (h *Handler) GetTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		res, err := h.uc.GetTask(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, err.Error())
			return

		}

		ctx.JSON(http.StatusOK, res)

	}
}

func (h *Handler) CreateTask() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		args := task.CreateTaskRequest{}

		err := ctx.ShouldBindJSON(&args)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		res := h.uc.CreateTask(ctx.Request.Context(), args)
		ctx.JSON(http.StatusCreated, res)

	}
}

func (h *Handler) GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := h.uc.GetUsers()

		ctx.JSON(http.StatusOK, res)
	}
}

func (h *Handler) GetTasks() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		res := h.uc.GetTasks()

		ctx.JSON(http.StatusOK, res)

	}
}

type useCase interface {
	GetTasks() []*task.Task
	CreateTask(context context.Context, args task.CreateTaskRequest) task.Task
	GetTask(id int) (task.Task, error)
	UpdateTask(context context.Context, args task.UpdateTaskRequest, id int) (task.Task, error)
	DeleteTask(context context.Context, id int) error
	GetUsers() []*user.User
	GetUser(id int) (user.User, error)
	CreateUser(ontext context.Context, args user.CreateUserRequest) user.User
	UpdateUser(context context.Context, args user.UpdateUserRequest, id int) (user.User, error)
	DeleteUser(context context.Context, id int) error
}

func New(uc useCase) *Handler {
	return &Handler{
		uc: uc,
	}
}
