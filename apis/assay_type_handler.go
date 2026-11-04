package apis

import (
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/protengplus/proteng-bff/configs"
)

func conductorUrlWithUser(c *gin.Context, path string) string {
	query := c.Request.URL.Query()
	query.Set("user_id", c.GetString("userId"))
	return configs.Config.ConductorUrl + path + "?" + query.Encode()
}

func userBody(c *gin.Context) map[string]interface{} {
	return map[string]interface{}{"user_id": c.GetString("userId")}
}

// @Summary Get all assay types
// @Description Retrieve the assay types of the current user with job and result counts
// @Tags Assay Types
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 401 {object} models.HttpResponseError "Unauthorized"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/assay-types [get]
func GetAllAssayTypes(c *gin.Context) {
	ForwardStrict[interface{}, interface{}](conductorUrlWithUser(c, "/assay-types"))(c)
}

// @Summary Create an assay type
// @Description Create a reusable assay score type
// @Tags Assay Types
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param assayType body models.AssayTypeInput true "Assay type object"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 409 {object} models.HttpResponseError "Name already exists"
// @Failure 500 {object} models.HttpResponseError "Internal Server Error"
// @Router /proteng-conductor/assay-types [post]
func CreateAssayType(c *gin.Context) {
	ForwardAddBody(configs.Config.ConductorUrl+"/assay-types", userBody(c))(c)
}

// @Summary Get an assay type
// @Description Retrieve one assay type of the current user
// @Tags Assay Types
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Assay type ID"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 403 {object} models.HttpResponseError "Forbidden"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Router /proteng-conductor/assay-types/{id} [get]
func GetAssayType(c *gin.Context) {
	ForwardStrict[interface{}, interface{}](conductorUrlWithUser(c, "/assay-types/"+url.PathEscape(c.Param("id"))))(c)
}

// @Summary Update an assay type
// @Description Rename an assay type or change its description
// @Tags Assay Types
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "Assay type ID"
// @Param assayType body models.AssayTypeInput true "Assay type object"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 403 {object} models.HttpResponseError "Forbidden"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 409 {object} models.HttpResponseError "Name already exists"
// @Router /proteng-conductor/assay-types/{id} [put]
func UpdateAssayType(c *gin.Context) {
	ForwardAddBody(configs.Config.ConductorUrl+"/assay-types/"+url.PathEscape(c.Param("id")), userBody(c))(c)
}

// @Summary Delete an assay type
// @Description Delete an assay type that no job or result uses
// @Tags Assay Types
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Assay type ID"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 403 {object} models.HttpResponseError "Forbidden"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Failure 409 {object} models.HttpResponseError "Assay type is in use"
// @Router /proteng-conductor/assay-types/{id} [delete]
func DeleteAssayType(c *gin.Context) {
	ForwardStrict[interface{}, interface{}](conductorUrlWithUser(c, "/assay-types/"+url.PathEscape(c.Param("id"))))(c)
}

// @Summary List library jobs
// @Description Jobs that have wet-lab results of this assay type, newest first
// @Tags Assay Types
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Assay type ID"
// @Param q query string false "Filter by job name"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 403 {object} models.HttpResponseError "Forbidden"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Router /proteng-conductor/assay-types/{id}/library/jobs [get]
func GetLibraryJobs(c *gin.Context) {
	ForwardStrict[interface{}, interface{}](conductorUrlWithUser(c, "/assay-types/"+url.PathEscape(c.Param("id"))+"/library/jobs"))(c)
}

// @Summary List library results of a job
// @Description Deduplicated wet-lab results of one job for this assay type
// @Tags Assay Types
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Assay type ID"
// @Param job_id query string true "Job ID"
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Failure 400 {object} models.HttpResponseError "Bad request"
// @Failure 403 {object} models.HttpResponseError "Forbidden"
// @Failure 404 {object} models.HttpResponseError "Not found"
// @Router /proteng-conductor/assay-types/{id}/library [get]
func GetLibraryResults(c *gin.Context) {
	ForwardStrict[interface{}, interface{}](conductorUrlWithUser(c, "/assay-types/"+url.PathEscape(c.Param("id"))+"/library"))(c)
}

// @Summary Export the library as csv
// @Description Headerless sequence,score csv of all or the given library results
// @Tags Assay Types
// @Security ApiKeyAuth
// @Produce text/csv
// @Param id path string true "Assay type ID"
// @Param ids query string false "Comma separated result IDs"
// @Success 200 {string} string "csv file"
// @Failure 403 {object} models.HttpResponseError "Forbidden"
// @Failure 404 {object} models.HttpResponseError "Nothing to export"
// @Router /proteng-conductor/assay-types/{id}/library/export [get]
func ExportLibrary(c *gin.Context) {
	ForwardWithRawDataResponse(conductorUrlWithUser(c, "/assay-types/"+url.PathEscape(c.Param("id"))+"/library/export"))(c)
}

// @Summary Export a job's wet-lab results as csv
// @Description Headerless sequence,score csv that the Upload Lab Input step accepts
// @Tags Experimental Results
// @Security ApiKeyAuth
// @Produce text/csv
// @Param job_id query string true "Job ID"
// @Success 200 {string} string "csv file"
// @Failure 403 {object} models.HttpResponseError "Forbidden"
// @Failure 404 {object} models.HttpResponseError "Nothing to export"
// @Router /proteng-conductor/mutations/experimental-results/export [get]
func ExportExperimentalResults(c *gin.Context) {
	ForwardWithRawDataResponse(conductorUrlWithUser(c, "/mutations/experimental-results/export"))(c)
}
