package controllers

import (
	"net/http"

	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/config"
	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/models"
	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/services"

	"github.com/labstack/echo"
)

// get all users
// Controller untuk mendapatkan semua pengguna
func GetAllUsersController(c echo.Context) error {
    userResponses, err := services.GetAllUsersService()
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "status":  http.StatusOK,
        "message": "success get all users",
        "users":   userResponses,
    })
}

// get user by id
func GetUserController(c echo.Context) error {
    userID := c.Param("id")

    userResponse, err := services.GetUserByIDService(userID)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "status": http.StatusOK,
        "message": "success get user by id",
        "user": userResponse,
    })
}
// create new user
func CreateUserController(c echo.Context) error {
    user := models.User{}
    c.Bind(&user)

    userResponse, err := services.CreateUserService(user)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "status":  http.StatusOK,
        "message": "success create new user",
        "user":    userResponse,
    })
}
// delete user by id
// Delete user by ID
func DeleteUserController(c echo.Context) error {
    userID := c.Param("id")

    if err := services.DeleteUserService(userID); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "status":  http.StatusOK,
        "message": "success delete user by id",
    })
}

// Update user by ID
func UpdateUserController(c echo.Context) error {
    user := models.User{}
    userID := c.Param("id")

    if err := config.DB.First(&user, userID).Error; err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    if err := c.Bind(&user); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    updatedUser, err := services.UpdateUserService(userID, user)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "status":  http.StatusOK,
        "message": "success update user by id",
        "user":    updatedUser,
    })
}
func LoginUserController(c echo.Context) error {
    user := models.User{}
    c.Bind(&user)
    email := user.Email
    password := user.Password

    userResponse, err := services.LoginUserService(email, password)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]interface{}{
            "message": "fail login",
            "error":   "invalid email or password",
        })
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "status":  http.StatusOK,
        "message": "success login",
        "user":    userResponse,
    })
}

