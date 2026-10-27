package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/protengplus/proteng-bff/configs"
)

// @Summary Get evaluation plugins
// @Description Retrieve the names of the evaluation plugins a job can use
// @Tags Evaluations
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/evaluation-plugins [get]
func GetEvaluationPlugins(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	Forward(conductorUrl + "/evaluation-plugins")(c)
}

// @Summary Get the evaluation runs of a job
// @Description Retrieve the evaluation runs of a job of the current user, newest first
// @Tags Evaluations
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Job ID"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 403 {object} models.HttpResponseError "Forbidden"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/jobs/{id}/evaluations [get]
func GetJobEvaluations(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	ForwardAddParam(conductorUrl+"/jobs/"+c.Param("id")+"/evaluations", map[string]interface{}{"user_id": c.GetString("userId")})(c)
}

// @Summary Rerun the evaluation of a job
// @Description Evaluate a completed mutation run of a completed job again with the given plugins. Without mutation_id the latest completed mutation run is used.
// @Tags Evaluations
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "Job ID"
// @Param evaluation body models.RerunEvaluationInput true "Rerun object"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 403 {object} models.HttpResponseError "Forbidden"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/jobs/{id}/evaluations [post]
func RerunEvaluation(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	ForwardAddBody(conductorUrl+"/jobs/"+c.Param("id")+"/evaluations", map[string]interface{}{"user_id": c.GetString("userId")})(c)
}

// @Summary Get an evaluation run
// @Description Retrieve an evaluation run of the current user with its results grouped by mutant
// @Tags Evaluations
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "EvaluationRun ID"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 403 {object} models.HttpResponseError "Forbidden"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/evaluations/{id} [get]
func GetEvaluation(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	ForwardAddParam(conductorUrl+"/evaluations/"+c.Param("id"), map[string]interface{}{"user_id": c.GetString("userId")})(c)
}
