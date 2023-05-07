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
	eJWT.GET("participants", controllers.GetParticipantsController)
	eJWT.GET("participants/:id", controllers.GetParticipantController)
	eJWT.POST("participants", controllers.CreateParticipantController)
	eJWT.DELETE("participants/:id", controllers.DeleteParticipantController)
	eJWT.PUT("participants/:id", controllers.UpdateParticipantController)
	// Registrations
	e.GET("registrations", controllers.GetRegistrationsController)
	e.GET("registrations/:id", controllers.GetRegistrationController)
	e.POST("registrations", controllers.CreateRegistrationController)
	e.DELETE("registrations/:id", controllers.DeleteRegistrationController)
	e.PUT("registrations/:id", controllers.UpdateRegistrationController)
	// Monitorings
	e.GET("monitorings", controllers.GetMonitoringsController)
	e.GET("monitorings/:exam_reg", controllers.GetMonitoringController)
	e.POST("monitorings", controllers.CreateMonitoringController)
	// e.DELETE("registrations/:id", controllers.DeleteRegistrationController)
	// e.PUT("registrations/:id", controllers.UpdateRegistrationController)


	return e
}
