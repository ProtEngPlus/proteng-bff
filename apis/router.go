package apis

import (
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
	userRouter.PATCH("/auth/changepassword", middleware.Authenticate(), ChangePassword)
	userRouter.POST("/auth/sendverification", SendVerification)
	userRouter.POST("/auth/verifyemail/:verificationToken", VerifyEmail)

	userRouter.GET("/admins", middleware.Authenticate(), middleware.Authorize("admin"), GetAllAdmins)
	userRouter.GET("/admins/:id", middleware.Authenticate(), middleware.Authorize("admin"), GetAdminByID)
	userRouter.POST("/admins", CreateAdmin)
	userRouter.PUT("/admins/:id", middleware.Authenticate(), middleware.Authorize("admin"), UpdateAdmin)
	userRouter.DELETE("/admins/:id", middleware.Authenticate(), middleware.Authorize("admin"), DeleteAdmin)
	userRouter.POST("/auth/login/admin", SignInAdmin)

	conductorRouter := r.Group("/proteng-conductor")
	// conductorUrl := configs.Config.ConductorUrl

	conductorRouter.GET("/jobs", middleware.Authenticate(), middleware.Authorize("user"), GetAllJobs)
	conductorRouter.GET("/jobs/:id", middleware.Authenticate(), middleware.Authorize("user"), GetJob)
	conductorRouter.POST("/jobs", middleware.Authenticate(), middleware.Authorize("user"), CreateJob)
	conductorRouter.PUT("/jobs/:id", middleware.Authenticate(), middleware.Authorize("user", "staff"), UpdateJob)
	conductorRouter.DELETE("/jobs/:id", middleware.Authenticate(), middleware.Authorize("user"), DeleteJob)
	conductorRouter.POST("/jobs/:id/run", middleware.Authenticate(), middleware.Authorize("user"), RunJob)
	conductorRouter.POST("/jobs/:id/:stage", middleware.Authenticate(), middleware.Authorize("user"), CreateDuplicateJob)

	conductorRouter.GET("/mutations", middleware.Authenticate(), middleware.Authorize("user"), GetAllMutations)
	conductorRouter.GET("/mutations/:id", middleware.Authenticate(), middleware.Authorize("user"), GetMutation)
	conductorRouter.POST("/mutations", middleware.Authenticate(), middleware.Authorize("user"), CreateMutation)
	conductorRouter.PUT("/mutations/:id", middleware.Authenticate(), middleware.Authorize("user"), UpdateMutation)
	conductorRouter.DELETE("/mutations/:id", middleware.Authenticate(), middleware.Authorize("user"), DeleteMutation)
	conductorRouter.POST("/mutations/:id/run", middleware.Authenticate(), middleware.Authorize("user"), RunMutation)

	conductorRouter.GET("/artifact/:bucketName/:objectName", middleware.Authenticate(), middleware.Authorize("user"), DownloadArtifact)

	conductorRouter.GET("/uniProt/:uniProtId", middleware.Authenticate(), middleware.Authorize("user"), GetProteinSequenceFromId)
	
	conductorRouter.GET("/query_results", middleware.Authenticate(), middleware.Authorize("user"), GetAllQueryResult)
	conductorRouter.GET("/query_results/:id", middleware.Authenticate(), middleware.Authorize("user"), GetQueryResult)
	conductorRouter.PUT("/query_results/:id", middleware.Authenticate(), middleware.Authorize("user"), UpdateQueryResult)
}
