package routes

import (
	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/constants"
	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/controllers"
	m "github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/middlewares"

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
	
	ClientJWT := e.Group("/")
	ClientJWT.Use(mid.JWT([]byte(constants.CLIEN_SECRET_JWT)))
	// e.GET("/users", controllers.GetUsersController)
	eJWT.GET("users", controllers.GetAllUsersController)
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
	e.POST("participant/login", controllers.LoginClientController)
	// Registrations
	eJWT.GET("registrations", controllers.GetRegistrationsController)
	eJWT.GET("registrations/:id", controllers.GetRegistrationController)
	eJWT.POST("registrations", controllers.CreateRegistrationController)
	eJWT.DELETE("registrations/:id", controllers.DeleteRegistrationController)
	eJWT.PUT("registrations/:id", controllers.UpdateRegistrationController)
	// Monitorings
	eJWT.GET("monitorings", controllers.GetMonitoringsController)
	// eJWT.GET("monitorings/:exam_reg", controllers.GetMonitoringController)
	ClientJWT.POST("monitorings", controllers.CreateMonitoringController)
	ClientJWT.POST("monitorings2", controllers.CreateMonitoringController2)
	// e.DELETE("registrations/:id", controllers.DeleteRegistrationController)
	// e.PUT("registrations/:id", controllers.UpdateRegistrationController)



	return e
}
