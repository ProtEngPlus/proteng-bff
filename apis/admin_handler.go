package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/protengplus/proteng-bff/configs"
)

// GetAllAdmins retrieves all admins
// @Summary Retrieve all admins
// @Description Retrieve a list of all admins
// @Tags Admin
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-user-mgmt/admins [get]
func GetAllAdmins(c *gin.Context) {
	usermgmtUrl := configs.Config.UserMgmtUrl
	Forward(usermgmtUrl + "/admins")(c)
}

// GetAdminByID retrieves an admin by ID
// @Summary Retrieve an admin by ID
// @Description Retrieve an admin by providing its ID
// @Tags Admin
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Admin ID"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-user-mgmt/admins/{id} [get]
func GetAdminByID(c *gin.Context) {
	// Forward the request to the appropriate endpoint
	usermgmtUrl := configs.Config.UserMgmtUrl
	Forward(usermgmtUrl + "/admins/" + c.Param("id"))(c)
}

// CreateAdmin creates a new admin
// @Summary Create a new admin
// @Description Creates a new admin with the provided details
// @Tags Admin
// @Accept json
// @Produce json
// @Param admin body models.AdminInput true "Admin object"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-user-mgmt/admins [post]
func CreateAdmin(c *gin.Context) {
	usermgmtUrl := configs.Config.UserMgmtUrl
	Forward(usermgmtUrl + "/admins")(c)
}

// UpdateAdmin updates an existing admin
// @Summary Update an existing admin
// @Description Update an existing admin with the provided details
// @Tags Admin
// @Accept json
// @Produce json
// @Param id path string true "Admin ID"
// @Param admin body models.AdminInput true "Admin object"
// @Security ApiKeyAuth
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-user-mgmt/admins/{id} [put]
func UpdateAdmin(c *gin.Context) {
	usermgmtUrl := configs.Config.UserMgmtUrl
	Forward(usermgmtUrl + "/admins/" + c.Param("id"))(c)
}

// DeleteAdmin deletes an admin by ID
// @Summary Delete an admin
// @Description Delete an admin by providing its ID
// @Tags Admin
// @Produce json
// @Param id path string true "Admin ID"
// @Security ApiKeyAuth
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-user-mgmt/admins/{id} [delete]
func DeleteAdmin(c *gin.Context) {
	usermgmtUrl := configs.Config.UserMgmtUrl
	Forward(usermgmtUrl + "/admins/" + c.Param("id"))(c)
}

// SignInAdmin signs in an admin by validating credentials
// @Summary Sign in an admin
// @Description Signs in an admin by validating the provided credentials
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body models.SignInAdminInput true "Admin credentials"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Router /proteng-user-mgmt/auth/login/admin [post]
func SignInAdmin(c *gin.Context) {
	usermgmtUrl := configs.Config.UserMgmtUrl
	Forward(usermgmtUrl + "/auth/login/admin")(c)
}
