package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"todo/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(context *gin.Context) {
	version := context.GetHeader("App-version")

	fmt.Println("app version from mobile: ", version)

	page := context.DefaultQuery("page", "1")           // Default to "1" if not provided
	pageSize := context.DefaultQuery("page_size", "10") // Default to "10" if not provided

	// Convert string values to int64 (or you can use int)
	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid page number"})
		return
	}

	pageIntSize, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid page size"})
		return
	}

	todos, err := models.GetAllTodos(pageInt, pageIntSize)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"success": true, "data": todos, "size": len(*todos)})
}

func AddTodo(context *gin.Context) {
	var todo models.Todo
	err := context.ShouldBindJSON(&todo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	err = models.AddTodos(todo)

	if err != nil {
		context.JSON(http.StatusCreated, gin.H{"message": "Cannot add todo", "success": false})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Todo created successfully", "success": true})
}

func GetTodoById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid todo ID", "error": err.Error()})
		return
	}

	todo, err := models.GetTodoById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Todo not found", "id": id, "success": false, "data": nil})
		return
	}

	response := gin.H{"success": true, "data": todo}
	context.JSON(http.StatusOK, response)
}

func DeleteTodoById(context *gin.Context) {
	todoId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id.",
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	todo := models.Todo{ID: todoId}
	err = todo.DeleteTodoById()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "success": false})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully", "todo": todo, "success": true})
}

func DeleteAll(context *gin.Context) {
	err := models.DeleteAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "success": false})
		return
	}
	context.JSON(http.StatusOK, gin.H{"success": true, "message": "All todos are deleted successfully"})
}

func UpdateTodo(context *gin.Context) {
	todoId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse todo id.", "success": false})
		return
	}

	var updateTodo models.Todo
	err = context.ShouldBindJSON(&updateTodo)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "success": false})
		return
	}

	updateTodo.ID = todoId
	err = updateTodo.UpdateTodo()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "success": false})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully", "success": true})
}
