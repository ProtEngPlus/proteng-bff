package apis

import (
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	userRouter := r.Group("/proteng-user-mgmt")
	usermgmtUrl := os.Getenv("USER_MGMT_URL")
	userRouter.GET("/users", Forward(usermgmtUrl+"/users"))
	userRouter.GET("/users/:id", func(c *gin.Context) { Forward(usermgmtUrl + "/users/" + c.Param("id"))(c) })
	userRouter.POST("/users", Forward(usermgmtUrl+"/users"))
	userRouter.PUT("/users/:id", func(c *gin.Context) { Forward(usermgmtUrl + "/users/" + c.Param("id"))(c) })
	userRouter.DELETE("/users/:id", func(c *gin.Context) { Forward(usermgmtUrl + "/users/" + c.Param("id"))(c) })
	userRouter.POST("/auth/login", Forward(usermgmtUrl+"/auth/login"))

	r.GET("/cats", ForwardNoStrict("https://catfact.ninja/fact"))
}
