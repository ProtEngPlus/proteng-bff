package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/protengplus/proteng-bff/configs"
)

// @Summary Get Protein Sequence From Id
// @Description Get Protein Sequence From UniProtId
// @Tags UniProt
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param uniProtId path string true "UniProt Id"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/uniProt/{uniProtId} [get]
func GetProteinSequenceFromId(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	Forward(conductorUrl + "/uniProt/" + c.Param("uniProtId"))(c)
}
