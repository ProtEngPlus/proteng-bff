package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/protengplus/proteng-bff/configs"
)

// GetAllUsers retrieves all users
// @Summary Retrieve all users
// @Description Retrieve a list of all users
// @Tags User management
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
// @Tags User management
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

// RegisterUser registers a new user and sends the verification email
// @Summary Register a new user
// @Description Creates a new user and sends a verification email in a single call
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body models.UserInput true "User object"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 409 {object} models.HttpResponseError "Email already registered"
// @Failure 502 {object} models.HttpResponseError "Bad Gateway"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-user-mgmt/auth/register [post]
func RegisterUser(c *gin.Context) {
	usermgmtUrl := configs.Config.UserMgmtUrl
	Forward(usermgmtUrl + "/auth/register")(c)
}

// UpdateUser updates an existing user
// @Summary Update an existing user
// @Description Update an existing user with the provided details
// @Tags User management
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
// @Tags User management
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
// @Tags Authentication
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
// @Tags User management
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
// @Tags User management
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
// @Tags User management
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

// ForgotPassword initiates the forgot password process by sending a reset email to the user
// @Summary Initiate forgot password process
// @Description Initiates the forgot password process by sending a reset email to the user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param body body models.ForgotPasswordInput true "User email"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 403 {object} models.HttpResponseError "Forbidden"
// @Failure 502 {object} models.HttpResponseError "Bad Gateway"
// @Router /proteng-user-mgmt/auth/forgotpassword [post]
func ForgotPassword(c *gin.Context) {
	usermgmtUrl := configs.Config.UserMgmtUrl
	Forward(usermgmtUrl + "/auth/forgotpassword")(c)
}

// ResetPassword resets the password using the provided reset token
// @Summary Reset password
// @Description Resets the password using the provided reset token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param resetToken path string true "Reset Token"
// @Param body body models.ResetPasswordInput true "New password"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 403 {object} models.HttpResponseError "Forbidden"
// @Router /proteng-user-mgmt/auth/resetpassword/{resetToken} [patch]
func ResetPassword(c *gin.Context) {
	usermgmtUrl := configs.Config.UserMgmtUrl
	resetToken := c.Param("resetToken")
	Forward(usermgmtUrl + "/auth/resetpassword/" + resetToken)(c)
}

// ChangePassword changes the password while logged in
// @Summary change password
// @Description Changes the password while logged in
// @Tags Authentication
// @Accept json
// @Produce json
// @Param body body models.ChangePasswordInput true "User credentials"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 403 {object} models.HttpResponseError "Forbidden"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Router /proteng-user-mgmt/auth/changepassword [patch]
func ChangePassword(c *gin.Context) {
	usermgmtUrl := configs.Config.UserMgmtUrl
	userID := c.GetString("userId")
	Forward(usermgmtUrl + "/auth/changepassword/" + userID)(c)
}

// SendVerification initiates the email verification process by sending a verification email to the user
// @Summary initiates verification process
// @Description Initiates the email verification process by sending a verification email to the user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param body body models.SendVerificationInput true "User email"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 403 {object} models.HttpResponseError "Forbidden"
// @Failure 502 {object} models.HttpResponseError "Bad Gateway"
// @Router /proteng-user-mgmt/auth/sentverification [post]
func SendVerification(c *gin.Context) {
	usermgmtUrl := configs.Config.UserMgmtUrl
	Forward(usermgmtUrl + "/auth/sentverification")(c)
}

// VerifyEmail verifies the email using the provided verification token
// @Summary Verify email
// @Description Verifies the email using the provided verification token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param verificationToken path string true "Verification Token"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 403 {object} models.HttpResponseError "Forbidden"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-user-mgmt/auth/verifyemail/{verificationToken} [patch]
func VerifyEmail(c *gin.Context) {
	usermgmtUrl := configs.Config.UserMgmtUrl
	verificationToken := c.Param("verificationToken")
	Forward(usermgmtUrl + "/auth/verifyemail/" + verificationToken)(c)
}
