package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/protengplus/proteng-bff/configs"
)

// @Summary Get all query results
// @Description Retrieve a list of all query results
// @Tags Query results
// @Security ApiKeyAuth
// @Produce json
// @Param job_id query string true "Filter by job ID"
// @Param is_selected query bool false "Filter by selected"
// @Param organism query string false "Filter by organism"
// @Param percentIdentityFrom query float64 false "Filter by percentIdentityFrom"
// @Param percentIdentityTo query float64 false "Filter by percentIdentityTo"
// @Param eValuesFrom query float64 false "Filter by eValuesFrom"
// @Param eValuesTo query float64 false "Filter by eValuesTo"
// @Param queryCoverFrom query float64 false "Filter by queryCoverFrom"
// @Param queryCoverTo query float64 false "Filter by queryCoverTo"
// @Param sort query string false "Sort by"
// @Param order query string false "Sort direction (asc/desc)"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/query_results [get]
func GetAllQueryResult(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	Forward(conductorUrl + "/query_results?" + c.Request.URL.RawQuery)(c)
}

// @Summary Get a query result by ID
// @Description Retrieve details of a query result by its ID
// @Tags Query results
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Query result ID"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/query_results/{id} [get]
func GetQueryResult(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	Forward(conductorUrl + "/query_results/" + c.Param("id"))(c)
}

// @Summary Update an existing query result
// @Description Update an existing query result with the provided details
// @Tags Query results
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "Query Result ID"
// @Param mutation body models.QueryResultInput true "Query result object"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/query_results/{id} [put]
func UpdateQueryResult(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	Forward(conductorUrl + "/query_results/" + c.Param("id"))(c)
}

// @Summary Download a query result by Job ID
// @Description Retrieve details of a query result by its Job ID
// @Tags Query results
// @Security ApiKeyAuth
// @Param id path string true "Query result Job ID"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/query_results/{id}/download [get]
func DownloadQueryResult(c *gin.Context) {
	conductorUrl := configs.Config.ConductorUrl
	ForwardWithRawDataResponse(conductorUrl + "/query_results/" + c.Param("id") + "/download")(c)
}
