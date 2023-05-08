package services

import (
	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/config"
	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/helpers"
	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/middlewares"
	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/models"
)

// Service untuk mendapatkan semua pengguna
func GetAllUsersService() ([]models.UserResponse, error) {
    users := []models.User{}
    if err := config.DB.Find(&users).Error; err != nil {
        return nil, err
    }

    userResponses := make([]models.UserResponse, len(users))
    for i, user := range users {
        userResponses[i] = models.UserResponse{int(user.ID), user.Name, user.Email, user.Phone}
    }

    return userResponses, nil
}

// Service untuk mendapatkan pengguna berdasarkan ID
func GetUserByIDService(userID string) (models.UserResponse, error) {
    user := models.User{}
    if err := config.DB.First(&user, userID).Error; err != nil {
        return models.UserResponse{}, err
    }

    userResponse := models.UserResponse{int(user.ID), user.Name, user.Email, user.Phone}
    return userResponse, nil
}

// Service untuk membuat pengguna baru
func CreateUserService(user models.User) (models.UserResponse, error) {
    // Hash password sebelum disimpan ke database
    hashedPassword, err := helpers.HashPassword(user.Password)
    if err != nil {
        return models.UserResponse{}, err
    }
    user.Password = string(hashedPassword)

    if err := config.DB.Save(&user).Error; err != nil {
        return models.UserResponse{}, err
    }

    userResponse := models.UserResponse{int(user.ID), user.Name, user.Email, user.Phone}
    return userResponse, nil
}
// Delete user by ID
func DeleteUserService(userID string) error {
    user := models.User{}

    if err := config.DB.First(&user, userID).Error; err != nil {
        return err
    }

    if err := config.DB.Delete(&user).Error; err != nil {
        return err
    }

    return nil
}

// Update user by ID
func UpdateUserService(userID string, updateUser models.User) (models.UserResponse, error) {
    user := models.User{}

    if err := config.DB.First(&user, userID).Error; err != nil {
        return models.UserResponse{}, err
    }

    // Hash password before saving to database
    if updateUser.Password != "" {
        hashedPassword, err := helpers.HashPassword(updateUser.Password)
        if err != nil {
            return models.UserResponse{}, err
        }
        updateUser.Password = string(hashedPassword)
    }

    if err := config.DB.Model(&user).Updates(updateUser).Error; err != nil {
        return models.UserResponse{}, err
    }

    updatedUser := models.User{}
    if err := config.DB.First(&updatedUser, userID).Error; err != nil {
        return models.UserResponse{}, err
    }

    userResponse := models.UserResponse{int(updatedUser.ID), updatedUser.Name, updatedUser.Email, updatedUser.Phone}
    return userResponse, nil
}
func LoginUserService(email, password string) (models.UserLoginResponse, error) {
    user := models.User{}
    // Get user from database by email
    if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
        return models.UserLoginResponse{}, err
    }

    // Verify password using bcrypt
    if err := helpers.CheckPasswordHash(password, user.Password); err != nil {
        return models.UserLoginResponse{}, err
    }

    // Generate JWT token
    token, err := middlewares.CreateToken(int(user.ID), user.Name)
    if err != nil {
        return models.UserLoginResponse{}, err
    }

    // Return user data and JWT token
    userResponse := models.UserLoginResponse{int(user.ID), user.Name, user.Email, user.Phone, token}
    return userResponse, nil
}


