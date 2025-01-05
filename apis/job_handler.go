package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/protengplus/proteng-bff/configs"
)

// @Summary Get all jobs
// @Description Retrieve a list of all jobs
// @Tags Jobs
// @Security ApiKeyAuth
// @Produce json
// @Param state query string false "Filter state of the jobs"
// @Param name query string false "Filter name of the job"
// @Param sort query string false "Sort by"
// @Param order query string false "Sort direction (asc/desc)"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/jobs [get]
func GetAllJobs(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	ForwardAddParam(conductorUrl+"/jobs?"+c.Request.URL.RawQuery, map[string]interface{}{"user_id": c.GetString("userId")})(c)
}

// @Summary Get job dashboard
// @Description Retrieve information about the jobs to show in the dashboard page
// @Tags Jobs
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/jobs/dashboard [get]
func GetJobDashboard(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	ForwardAddParam(conductorUrl+"/jobs/dashboard", map[string]interface{}{"user_id": c.GetString("userId")})(c)
}

// @Summary Get a job by ID
// @Description Retrieve details of a job by its ID
// @Tags Jobs
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Job ID"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/jobs/{id} [get]
func GetJob(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	Forward(conductorUrl + "/jobs/" + c.Param("id"))(c)
}

// @Summary Create a new job
// @Description Create a new job with the provided details
// @Tags Jobs
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param job body models.JobInput true "Job object"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/jobs [post]
func CreateJob(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	ForwardAddBody(conductorUrl+"/jobs", map[string]interface{}{"user_id": c.GetString("userId")})(c)
}

// @Summary Update an existing job
// @Description Update an existing job with the provided details
// @Tags Jobs
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "Job ID"
// @Param job body models.JobInput true "Job object"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/jobs/{id} [put]
func UpdateJob(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	Forward(conductorUrl + "/jobs/" + c.Param("id"))(c)
}

// @Summary Delete a job by ID
// @Description Delete a job by its ID
// @Tags Jobs
// @Security ApiKeyAuth
// @Param id path string true "Job ID"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/jobs/{id} [delete]
func DeleteJob(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	Forward(conductorUrl + "/jobs/" + c.Param("id"))(c)
}

// @Summary Start a job by ID
// @Description Start a job by its ID
// @Tags Jobs
// @Security ApiKeyAuth
// @Param id path string true "Job ID"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/jobs/{id}/run [post]
func RunJob(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	Forward(conductorUrl + "/jobs/" + c.Param("id") + "/run")(c)
}

// @Summary Create a duplicate job based on another job
// @Description Create a duplicate job based on another job with the provided details
// @Tags Jobs
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "Job ID"
// @Param stage path string true "Stage ID"
// @Param job body models.DuplicateJobInput true "Job object"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/jobs/{id}/{stage} [post]
func CreateDuplicateJob(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	Forward(conductorUrl + "/jobs/" + c.Param("id") + "/" + c.Param("stage"))(c)
}
