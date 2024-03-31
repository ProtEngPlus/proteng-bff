package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/protengplus/proteng-bff/configs"
)

// @Summary Download Artifact
// @Description Download an artifact of a job
// @Tags Artifact
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param bucketName path string true "Bucket Name"
// @Param objectName path string true "Object Name"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/artifact/{bucketName}/{objectName} [get]
func DownloadArtifact(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	Forward(conductorUrl + "/artifact/" + c.Param("bucketName") + "/" + c.Param("objectName"))(c)
}
