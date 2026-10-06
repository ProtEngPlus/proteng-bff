package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/protengplus/proteng-bff/configs"
)

// @Summary Get all experimental results
// @Description Retrieve the wet-lab results of the current user
// @Tags Experimental Results
// @Security ApiKeyAuth
// @Produce json
// @Param job_id query string false "Filter by job ID"
// @Param mutation_id query string false "Filter by mutation ID"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/mutations/experimental-results [get]
func GetAllExperimentalResults(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	ForwardAddParam(conductorUrl+"/mutations/experimental-results?"+c.Request.URL.RawQuery, map[string]interface{}{"user_id": c.GetString("userId")})(c)
}

// @Summary Create or overwrite an experimental result
// @Description Save the wet-lab result of a mutation result, replacing any existing one
// @Tags Experimental Results
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param mutationResultId path string true "MutationResult ID"
// @Param experimentalResult body models.ExperimentalResultInput true "ExperimentalResult object"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 403 {object} models.HttpResponseError "Forbidden"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/mutations/experimental-results/by-mutation-result/{mutationResultId} [put]
func UpsertExperimentalResult(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	ForwardAddBody(conductorUrl+"/mutations/experimental-results/by-mutation-result/"+c.Param("mutationResultId"), map[string]interface{}{"user_id": c.GetString("userId")})(c)
}

// @Summary Delete an experimental result
// @Description Clear the wet-lab result of a mutation result
// @Tags Experimental Results
// @Security ApiKeyAuth
// @Produce json
// @Param mutationResultId path string true "MutationResult ID"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 403 {object} models.HttpResponseError "Forbidden"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/mutations/experimental-results/by-mutation-result/{mutationResultId} [delete]
func DeleteExperimentalResult(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	ForwardAddParam(conductorUrl+"/mutations/experimental-results/by-mutation-result/"+c.Param("mutationResultId"), map[string]interface{}{"user_id": c.GetString("userId")})(c)
}
