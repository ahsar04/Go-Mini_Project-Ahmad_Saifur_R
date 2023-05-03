package routes

import (
	"code_structure/constants"
	"code_structure/controllers"
	m "code_structure/middlewares"

	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo{
	
	e := echo.New()
	// Route / to handler function
	// users
	m.LogMiddleware(e)
	
	eJWT := e.Group("/")
	eJWT.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	// e.GET("/users", controllers.GetUsersController)
	eJWT.GET("users", controllers.GetUsersController)
	eJWT.GET("users/:id", controllers.GetUserController)
	e.POST("users", controllers.CreateUserController)
	e.POST("users/login", controllers.LoginUserController)
	eJWT.DELETE("users/:id", controllers.DeleteUserController)
	eJWT.PUT("users/:id", controllers.UpdateUserController)
	// exams
	eJWT.GET("exams", controllers.GetExamsController)
	eJWT.GET("exams/:id", controllers.GetExamController)
	eJWT.POST("exams", controllers.CreateExamController)
	eJWT.DELETE("exams/:id", controllers.DeleteExamController)
	eJWT.PUT("exams/:id", controllers.UpdateExamController)
	// participans
	e.GET("/participants", controllers.GetParticipantsController)
	e.GET("/participants/:id", controllers.GetParticipantController)
	e.POST("/participants", controllers.CreateParticipantController)
	e.DELETE("/participants/:id", controllers.DeleteParticipantController)
	e.PUT("/participants/:id", controllers.UpdateParticipantController)


	return e
}
