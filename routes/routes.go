package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo/controller"
	"todo/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/index", func(c *gin.Context) { c.String(http.StatusOK, "Welcome to Todo app") })
	authenticated := server.Group("/auth")
	authenticated.Use(middlewares.AuthMiddleware())

	todoRoutes := authenticated.Group("/todo")
	todoRoutes.PUT("/todo:id", controller.UpdateTodo)
	todoRoutes.GET("/todoById/:id", controller.GetTodoById)
	todoRoutes.GET("/all", controller.GetTodos)
	todoRoutes.DELETE("/delete:id", controller.DeleteTodoById)
	todoRoutes.DELETE("/deleteAll", controller.DeleteAll)
	todoRoutes.POST("/addTodo", controller.AddTodo)
	server.POST("/login", controller.Login)
	server.POST("/signup", controller.Signup)
}
