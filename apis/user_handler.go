package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/protengplus/proteng-bff/configs"
)

// GetAllUsers retrieves all users
// @Summary Retrieve all users
// @Description Retrieve a list of all users
// @Tags user management
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-user-mgmt/users [get]
func GetAllUsers(c *gin.Context) {
	usermgmtUrl := configs.Config.UserMgmtUrl
	Forward(usermgmtUrl + "/users")(c)
}

// GetUserByID retrieves a user by ID
// @Summary Retrieve a user by ID
// @Description Retrieve a user by providing its ID
// @Tags user management
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-user-mgmt/users/{id} [get]
func GetUserByID(c *gin.Context) {
	// Forward the request to the appropriate endpoint
	usermgmtUrl := configs.Config.UserMgmtUrl
	Forward(usermgmtUrl + "/users/" + c.Param("id"))(c)
}

// CreateUser creates a new user
// @Summary Create a new user
// @Description Creates a new user with the provided details
// @Tags user management
// @Accept json
// @Produce json
// @Param user body models.UserInput true "User object"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-user-mgmt/users [post]
func CreateUser(c *gin.Context) {
	usermgmtUrl := configs.Config.UserMgmtUrl
	Forward(usermgmtUrl + "/users")(c)
}

// UpdateUser updates an existing user
// @Summary Update an existing user
// @Description Update an existing user with the provided details
// @Tags user management
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body models.UserInput true "User object"
// @Security ApiKeyAuth
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-user-mgmt/users/{id} [put]
func UpdateUser(c *gin.Context) {
	usermgmtUrl := configs.Config.UserMgmtUrl
	Forward(usermgmtUrl + "/users/" + c.Param("id"))(c)
}

// DeleteUser deletes a user by ID
// @Summary Delete a user
// @Description Delete a user by providing its ID
// @Tags user management
// @Produce json
// @Param id path string true "User ID"
// @Security ApiKeyAuth
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-user-mgmt/users/{id} [delete]
func DeleteUser(c *gin.Context) {
	usermgmtUrl := configs.Config.UserMgmtUrl
	Forward(usermgmtUrl + "/users/" + c.Param("id"))(c)
}

// SignInUser signs in a user by validating credentials
// @Summary Sign in a user
// @Description Signs in a user by validating the provided credentials
// @Tags authentication
// @Accept json
// @Produce json
// @Param credentials body models.SignInInput true "User credentials"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Router /proteng-user-mgmt/auth/login [post]
func SignInUser(c *gin.Context) {
	usermgmtUrl := configs.Config.UserMgmtUrl
	Forward(usermgmtUrl + "/auth/login")(c)
}

// @Summary Get current user
// @Description Retrieve details of the currently authenticated user
// @Tags user management
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-user-mgmt/me [get]
func GetMe(c *gin.Context) {
	usermgmtUrl := configs.Config.UserMgmtUrl
	userID := c.GetString("userId")
	Forward(usermgmtUrl + "/users/" + userID)(c)
}

// UpdateMe updates the currently authenticated user
// @Summary Update the currently authenticated user
// @Description Update details of the currently authenticated user
// @Tags user management
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body models.UserInput true "User object"
// @Security ApiKeyAuth
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-user-mgmt/me [put]
func UpdateMe(c *gin.Context) {
	usermgmtUrl := configs.Config.UserMgmtUrl
	userID := c.GetString("userId")
	Forward(usermgmtUrl + "/users/" + userID)(c)
}

// DeleteMe deletes the currently authenticated user
// @Summary Delete the currently authenticated user
// @Description Delete the currently authenticated user
// @Tags user management
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-user-mgmt/me [delete]
func DeleteMe(c *gin.Context) {
	usermgmtUrl := configs.Config.UserMgmtUrl
	userID := c.GetString("userId")
	Forward(usermgmtUrl + "/users/" + userID)(c)
}
