package apis

import (
	"github.com/protengplus/proteng-bff/configs"
	"github.com/protengplus/proteng-bff/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	userRouter := r.Group("/proteng-user-mgmt")
	// usermgmtUrl := configs.Config.UserMgmtUrl

	userRouter.GET("/users", middleware.Authenticate(), middleware.Authorize("admin"), GetAllUsers)
	userRouter.GET("/users/:id", middleware.Authenticate(), middleware.Authorize("admin"), GetUserByID)
	userRouter.POST("/users", CreateUser)
	userRouter.PUT("/users/:id", middleware.Authenticate(), middleware.Authorize("admin"), UpdateUser)
	userRouter.DELETE("/users/:id", middleware.Authenticate(), middleware.Authorize("admin"), DeleteUser)
	userRouter.POST("/auth/login", SignInUser)
	userRouter.GET("/me", middleware.Authenticate(), GetMe)
	userRouter.PUT("/me", middleware.Authenticate(), UpdateMe)
	userRouter.DELETE("/me", middleware.Authenticate(), DeleteMe)
	userRouter.POST("/auth/forgotpassword", ForgotPassword)
	userRouter.PATCH("/auth/resetpassword/:resetToken", ResetPassword)

	userRouter.GET("/admins", middleware.Authenticate(), middleware.Authorize("admin"), GetAllAdmins)
	userRouter.GET("/admins/:id", middleware.Authenticate(), middleware.Authorize("admin"), GetAdminByID)
	userRouter.POST("/admins", CreateAdmin)
	userRouter.PUT("/admins/:id", middleware.Authenticate(), middleware.Authorize("admin"), UpdateAdmin)
	userRouter.DELETE("/admins/:id", middleware.Authenticate(), middleware.Authorize("admin"), DeleteAdmin)
	userRouter.POST("/auth/login/admin", SignInAdmin)

	conductorRouter := r.Group("/proteng-conductor")
	conductorUrl := configs.Config.ConductorUrl

	conductorRouter.GET("/jobs", middleware.Authenticate(), middleware.Authorize("user"), GetAllJobs)
	conductorRouter.GET("/jobs/:id", middleware.Authenticate(), middleware.Authorize("user"), GetJob)
	conductorRouter.POST("/jobs", middleware.Authenticate(), middleware.Authorize("user"), CreateJob)
	conductorRouter.PUT("/jobs/:id", middleware.Authenticate(), middleware.Authorize("user", "staff"), UpdateJob)
	conductorRouter.DELETE("/jobs/:id", middleware.Authenticate(), middleware.Authorize("user"), DeleteJob)
	conductorRouter.POST("/jobs/:id/run", middleware.Authenticate(), middleware.Authorize("user"), RunJob)
	conductorRouter.POST("/jobs/:id/:stage", middleware.Authenticate(), middleware.Authorize("user"), CreateDuplicateJob)

	conductorRouter.GET("/mutations", middleware.Authenticate(), middleware.Authorize("user"), func(c *gin.Context) { Forward(conductorUrl + "/mutations?" + c.Request.URL.RawQuery)(c) })
	conductorRouter.GET("/mutations/:id", middleware.Authenticate(), middleware.Authorize("user"), func(c *gin.Context) { Forward(conductorUrl + "/mutations/" + c.Param("id"))(c) })
	conductorRouter.POST("/mutations", middleware.Authenticate(), middleware.Authorize("user"), Forward(conductorUrl+"/mutations"))
	conductorRouter.PUT("/mutations/:id", middleware.Authenticate(), middleware.Authorize("user"), func(c *gin.Context) { Forward(conductorUrl + "/mutations/" + c.Param("id"))(c) })
	conductorRouter.DELETE("/mutations/:id", middleware.Authenticate(), middleware.Authorize("user"), func(c *gin.Context) { Forward(conductorUrl + "/mutations/" + c.Param("id"))(c) })
	conductorRouter.POST("/mutations/:id/run", middleware.Authenticate(), middleware.Authorize("user"), func(c *gin.Context) { Forward(conductorUrl + "/mutations/" + c.Param("id") + "/run")(c) })
}
