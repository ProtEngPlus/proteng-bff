package apis

import (
	"os"

	"proteng-bff/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	userRouter := r.Group("/proteng-user-mgmt")
	usermgmtUrl := os.Getenv("USER_MGMT_URL")
	userRouter.GET("/users", Forward(usermgmtUrl+"/users"))
	userRouter.GET("/users/:id", func(c *gin.Context) { Forward(usermgmtUrl + "/users/" + c.Param("id"))(c) })
	userRouter.POST("/users", Forward(usermgmtUrl+"/users"))
	userRouter.PUT("/users/:id", middleware.AuthenticateUser(), func(c *gin.Context) { Forward(usermgmtUrl + "/users/" + c.Param("id"))(c) })
	userRouter.DELETE("/users/:id", func(c *gin.Context) { Forward(usermgmtUrl + "/users/" + c.Param("id"))(c) })
	userRouter.POST("/auth/login", Forward(usermgmtUrl+"/auth/login"))
	userRouter.GET("/me", middleware.AuthenticateUser(), func(c *gin.Context) { Forward(usermgmtUrl + "/users/" + c.GetString("userId"))(c) })

	conductorRouter := r.Group("/proteng-conductor")
	conductorUrl := os.Getenv("CONDUCTOR_URL")
	conductorRouter.GET("/jobs", Forward(conductorUrl+"/jobs"))
	conductorRouter.GET("/jobs/:id", func(c *gin.Context) { Forward(conductorUrl + "/jobs/" + c.Param("id"))(c) })
	conductorRouter.POST("/jobs", middleware.AuthenticateUser(), middleware.Authorize("user"), Forward(conductorUrl+"/jobs"))
	conductorRouter.PUT("/jobs/:id", func(c *gin.Context) { Forward(conductorUrl + "/jobs/" + c.Param("id"))(c) })
	conductorRouter.DELETE("/jobs/:id", func(c *gin.Context) { Forward(conductorUrl + "/jobs/" + c.Param("id"))(c) })
	conductorRouter.POST("/jobs/:id/run", func(c *gin.Context) { Forward(conductorUrl + "/jobs/" + c.Param("id") + "/run")(c) })
	conductorRouter.POST("/jobs/:id/:stage", func(c *gin.Context) { Forward(conductorUrl + "/jobs/" + c.Param("id") + "/" + c.Param("stage"))(c) })

	conductorRouter.GET("/mutations", Forward(conductorUrl+"/mutations"))
	conductorRouter.GET("/mutations/:id", func(c *gin.Context) { Forward(conductorUrl + "/mutations/" + c.Param("id"))(c) })
	conductorRouter.POST("/mutations", Forward(conductorUrl+"/mutations"))
	conductorRouter.PUT("/mutations/:id", func(c *gin.Context) { Forward(conductorUrl + "/mutations/" + c.Param("id"))(c) })
	conductorRouter.DELETE("/mutations/:id", func(c *gin.Context) { Forward(conductorUrl + "/mutations/" + c.Param("id"))(c) })

	r.GET("/cats", ForwardNoStrict("https://catfact.ninja/fact"))
}
