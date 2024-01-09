package apis

import (
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	userRouter := r.Group("/proteng-user-mgmt")
	usermgmtHost := os.Getenv("USER_MGMT_HOST")
	userRouter.GET("/users", Forward(usermgmtHost+"/users"))

	r.GET("/cats", ForwardNoStrict("https://catfact.ninja/fact"))
}
