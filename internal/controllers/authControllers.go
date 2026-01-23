package controllers

import (
  "os"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/surajit/notes-api/internal/models"
  "github.com/surajit/notes-api/internal/services"
  "github.com/surajit/notes-api/utils"
)

type AuthController struct {
  authService *services.AuthService
}

func NewAuthController() *AuthController {
  return &AuthController{
    authService: services.NewAuthService(),
  }
}

func (ac *AuthController) Register (c *gin.Context) {
  var signupDto models.SignupDTO

  // bind json data to signupDto and validate fields
  if err := c.ShouldBindJSON(&signupDto); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
      "error": err.Error(),
    })
    return
  }

  // call the service layer to create a new user
  user, err := ac.authService.RegisterUser(signupDto)
  if err != nil {
    c.JSON(http.StatusConflict, gin.H{
      "success": false,
      "error": err.Error(),
    })
    return
  }

  // generate jwt token
  token, err := utils.GenerateToken(user.ID.String(), os.Getenv("JWT_SECRET"))
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "success": false,
      "error": err.Error(),
    })
    return
  }

  // set token in cookie
  c.SetCookie("access_token", token, 3600, "/", "", true, true)

  // return success response
  c.JSON(http.StatusCreated, gin.H{
    "success": true,
    "message": "User created successfully",
    "user": user,
    "token": token,
  })
}

func (ac *AuthController) Login (c *gin.Context) {
  var loginDTO models.LoginDTO

  // bind json data to loginDTO and validate fields
  if err := c.ShouldBindJSON(&loginDTO); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
      "error": err.Error(),
    })
    return
  }

  // call the service layer to login the user
  user, err := ac.authService.LoginUser(loginDTO)
  if err != nil {
    c.JSON(http.StatusUnauthorized, gin.H{
      "success": false,
      "error": err.Error(),
    })
    return
  }

  // generate jwt token
  token, err := utils.GenerateToken(user.ID.String(), os.Getenv("JWT_SECRET"))
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "success": false,
      "error": err.Error(),
    })
    return
  }

  // set token in cookie
  c.SetCookie("access_token", token, 3600, "/", "", true, true)

  // return success response
  c.JSON(http.StatusOK, gin.H{
    "success": true,
    "message": "User logged in successfully",
    "user": user,
    "token": token,
  })
}

func (ac *AuthController) Logout (c *gin.Context) {
  // clear the cookie from the client
  c.SetCookie("access_token", "", -1, "/", "", true, true)
  
  c.JSON(http.StatusOK, gin.H{
    "success": true,
    "message": "User logged out successfully",
  })
}
