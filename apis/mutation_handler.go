package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/protengplus/proteng-bff/configs"
)

// @Summary Get all mutations
// @Description Retrieve a list of all mutations
// @Tags Mutations
// @Security ApiKeyAuth
// @Produce json
// @Param job_id query string true "Filter by job ID"
// @Param sort query string false "Sort by"
// @Param order query string false "Sort direction (asc/desc)"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/mutations [get]
func GetAllMutations(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	ForwardAddParam(conductorUrl+"/mutations?"+c.Request.URL.RawQuery, map[string]interface{}{"user_id": c.GetString("userId")})(c)
}

// @Summary Get a mutation by ID
// @Description Retrieve details of a mutation by its ID
// @Tags Mutations
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Mutation ID"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/mutations/{id} [get]
func GetMutation(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	Forward(conductorUrl + "/mutations/" + c.Param("id"))(c)
}

// @Summary Create a new mutation
// @Description Create a new mutation with the provided details
// @Tags Mutations
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param mutation body models.MutationInput true "Mutation object"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/mutations [post]
func CreateMutation(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	ForwardAddBody(conductorUrl+"/mutations", map[string]interface{}{"user_id": c.GetString("userId")})(c)
}

// @Summary Update an existing mutation
// @Description Update an existing mutation with the provided details
// @Tags Mutations
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "Mutation ID"
// @Param mutation body models.MutationInput true "Mutation object"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/mutations/{id} [put]
func UpdateMutation(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	Forward(conductorUrl + "/mutations/" + c.Param("id"))(c)
}

// @Summary Delete a mutation by ID
// @Description Delete a mutation by its ID
// @Tags Mutations
// @Security ApiKeyAuth
// @Param id path string true "Mutation ID"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/mutations/{id} [delete]
func DeleteMutation(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	Forward(conductorUrl + "/mutations/" + c.Param("id"))(c)
}

// @Summary Start a mutation by ID
// @Description Start a mutation by its ID
// @Tags Mutations
// @Security ApiKeyAuth
// @Param id path string true "Mutation ID"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/mutations/{id}/run [post]
func RunMutation(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	Forward(conductorUrl + "/mutations/" + c.Param("id") + "/run")(c)
}

// @Summary Download mutation results
// @Description Retrieve a list of all mutation results for a mutation and download them as a csv file
// @Tags Mutations
// @Security ApiKeyAuth
// @Param id path string true "Mutation ID"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/mutations/{id}/download [get]
func DownloadMutationResults(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	ForwardWithRawDataResponse(conductorUrl + "/mutations/" + c.Param("id") + "/download")(c)
}

// @Summary Get all mutation results
// @Description Retrieve a list of all mutation results for a mutation
// @Tags Mutations
// @Security ApiKeyAuth
// @Produce json
// @Param mutation_id query string true "Filter by mutation ID"
// @Param sort query string false "Sort by"
// @Param order query string false "Sort direction (asc/desc)"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/mutations/results [get]
func GetAllMutationResults(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	Forward(conductorUrl + "/mutations/results?" + c.Request.URL.RawQuery)(c)
}

// @Summary Update an existing mutation result
// @Description Update an existing mutation result with the provided details
// @Tags Mutation
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "MutationResult ID"
// @Param mutationResult body models.MutationResultInput true "MutationResult object"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/mutations/results/{resultId} [put]
func UpdateMutationResult(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	Forward(conductorUrl + "/mutations/results/" + c.Param("resultId"))(c)
}
