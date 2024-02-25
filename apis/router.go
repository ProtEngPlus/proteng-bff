package apis

import (
	"github.com/protengplus/proteng-bff/configs"
	"github.com/protengplus/proteng-bff/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	userRouter := r.Group("/proteng-user-mgmt")
	usermgmtUrl := configs.Config.UserMgmtUrl

	userRouter.GET("/users", middleware.Authenticate(), middleware.Authorize("admin"), GetAllUsers)
	userRouter.GET("/users/:id", middleware.Authenticate(), middleware.Authorize("admin"), GetUserByID)
	userRouter.POST("/users", CreateUser)
	userRouter.PUT("/users/:id", middleware.Authenticate(), middleware.Authorize("admin"), UpdateUser)
	userRouter.DELETE("/users/:id", middleware.Authenticate(), middleware.Authorize("admin"), DeleteUser)
	userRouter.POST("/auth/login", SignInUser)
	userRouter.GET("/me", middleware.Authenticate(), GetMe)
	userRouter.PUT("/me", middleware.Authenticate(), UpdateMe)
	userRouter.DELETE("/me", middleware.Authenticate(), DeleteMe)

	userRouter.GET("/admins", middleware.Authenticate(), middleware.Authorize("admin"), Forward(usermgmtUrl+"/admins"))
	userRouter.GET("/admins/:id", middleware.Authenticate(), middleware.Authorize("admin"), func(c *gin.Context) { Forward(usermgmtUrl + "/admins/" + c.Param("id"))(c) })
	userRouter.POST("/admins", Forward(usermgmtUrl+"/admins"))
	userRouter.PUT("/admins/:id", middleware.Authenticate(), middleware.Authorize("admin"), func(c *gin.Context) { Forward(usermgmtUrl + "/admins/" + c.Param("id"))(c) })
	userRouter.DELETE("/admins/:id", middleware.Authenticate(), middleware.Authorize("admin"), func(c *gin.Context) { Forward(usermgmtUrl + "/admins/" + c.Param("id"))(c) })
	userRouter.POST("/auth/login/admin", Forward(usermgmtUrl+"/auth/login/admin"))

	conductorRouter := r.Group("/proteng-conductor")
	conductorUrl := configs.Config.ConductorUrl

	conductorRouter.GET("/jobs", middleware.Authenticate(), middleware.Authorize("user"), func(c *gin.Context) {
		ForwardAddParam(conductorUrl+"/jobs?"+c.Request.URL.RawQuery, map[string]interface{}{"user_id": c.GetString("userId")})(c)
	})
	conductorRouter.GET("/jobs/:id", middleware.Authenticate(), middleware.Authorize("user"), func(c *gin.Context) { Forward(conductorUrl + "/jobs/" + c.Param("id"))(c) })
	conductorRouter.POST("/jobs", middleware.Authenticate(), middleware.Authorize("user"), func(c *gin.Context) {
		ForwardAddBody(conductorUrl+"/jobs", map[string]interface{}{"user_id": c.GetString("userId")})(c)
	})
	conductorRouter.PUT("/jobs/:id", middleware.Authenticate(), middleware.Authorize("user", "staff"), func(c *gin.Context) { Forward(conductorUrl + "/jobs/" + c.Param("id"))(c) })
	conductorRouter.DELETE("/jobs/:id", middleware.Authenticate(), middleware.Authorize("user"), func(c *gin.Context) { Forward(conductorUrl + "/jobs/" + c.Param("id"))(c) })
	conductorRouter.POST("/jobs/:id/run", middleware.Authenticate(), middleware.Authorize("user"), func(c *gin.Context) { Forward(conductorUrl + "/jobs/" + c.Param("id") + "/run")(c) })
	conductorRouter.POST("/jobs/:id/:stage", middleware.Authenticate(), middleware.Authorize("user"), func(c *gin.Context) { Forward(conductorUrl + "/jobs/" + c.Param("id") + "/" + c.Param("stage"))(c) })

	conductorRouter.GET("/mutations", middleware.Authenticate(), middleware.Authorize("user"), func(c *gin.Context) { Forward(conductorUrl + "/mutations?" + c.Request.URL.RawQuery)(c) })
	conductorRouter.GET("/mutations/:id", middleware.Authenticate(), middleware.Authorize("user"), func(c *gin.Context) { Forward(conductorUrl + "/mutations/" + c.Param("id"))(c) })
	conductorRouter.POST("/mutations", middleware.Authenticate(), middleware.Authorize("user"), Forward(conductorUrl+"/mutations"))
	conductorRouter.PUT("/mutations/:id", middleware.Authenticate(), middleware.Authorize("user"), func(c *gin.Context) { Forward(conductorUrl + "/mutations/" + c.Param("id"))(c) })
	conductorRouter.DELETE("/mutations/:id", middleware.Authenticate(), middleware.Authorize("user"), func(c *gin.Context) { Forward(conductorUrl + "/mutations/" + c.Param("id"))(c) })
}
