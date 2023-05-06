package controllers

import (
	"code_structure/config"
	"code_structure/helpers"
	"code_structure/middlewares"
	"code_structure/models"
	"net/http"

	"github.com/labstack/echo"
)

// get all users
func GetUsersController(c echo.Context) error {
	users:= models.User{}
	c.Bind(&users)


	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userResponse := models.UserResponse{int(users.ID),users.Name,users.Email,users.Phone}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success get all users",
		"users":   userResponse,
	})
}
// get user by id
func GetUserController(c echo.Context) error {
	user := models.User{}
	userID := c.Param("id")

	if err := config.DB.First(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userResponse := models.UserResponse{int(user.ID),user.Name,user.Email,user.Phone}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success get user by id",
		"user":    userResponse,
	})
}
// create new user
func CreateUserController(c echo.Context) error {
    user := models.User{}
    c.Bind(&user)

    // Hash password before saving to database
    hashedPassword, err := helpers.HashPassword(user.Password)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    user.Password = string(hashedPassword)

    if err := config.DB.Save(&user).Error; err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
	userResponse := models.UserResponse{int(user.ID),user.Name,user.Email,user.Phone}

    return c.JSON(http.StatusOK, map[string]interface{}{
        "status":  http.StatusOK,
        "message": "success create new user",
        "user":    userResponse,
    })
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	user := models.User{}
	userID := c.Param("id")

	if err := config.DB.First(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success delete user by id",
	})
}
// update user by id
func UpdateUserController(c echo.Context) error {
	user := models.User{}
	userID := c.Param("id")

	if err := config.DB.First(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success update user by id",
		"user":    user,
	})
}
func LoginUserController(c echo.Context) error {
    user := models.User{}
    c.Bind(&user)
	pass := user.Password

    // Get user from database by email
    if err := config.DB.Where("email = ?", user.Email).First(&user).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "message": "fail login",
            "error":   err.Error(),
        })
    }

    // Verify password using bcrypt
    if err := helpers.CheckPasswordHash(pass, user.Password); err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]interface{}{
            "message": "fail login",
            "error":   "invalid email or password",
        })
    }

    // Generate JWT token
    token, err := middlewares.CreateToken(int(user.ID), user.Name)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "message": "fail login",
            "error":   err.Error(),
        })
    }

    // Return user data and JWT token
    userResponse := models.UserLoginResponse{int(user.ID), user.Name, user.Email, user.Phone, token}
    return c.JSON(http.StatusOK, map[string]interface{}{
        "status":  http.StatusOK,
        "message": "success login",
        "user":    userResponse,
    })
}
