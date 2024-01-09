package apis

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"proteng-bff/models"
	"proteng-bff/utils/apiutil"
)

// Forward the request to another service with the common response format
func Forward(url string) func(c *gin.Context) {
	return ForwardStrict[models.HttpResponse, interface{}](url)
}

// ForwardNoStrict is Forward but the request and response body can be anything
func ForwardNoStrict(url string) func(c *gin.Context) {
	return ForwardStrict[interface{}, interface{}](url)
}

// ForwardStrict will have to specify the request and response struct
func ForwardStrict[Resp any, Req any](url string) func(c *gin.Context) {
	return func(gctx *gin.Context) {
		var body Req
		err := gctx.Bind(&body)
		if err != nil {
			apiutil.ApiResponseErrorBadRequest(gctx, err)
			return
		}

		var bodyBuffer []byte
		bodyBuffer, err = json.Marshal(body)
		if err != nil {
			apiutil.ApiResponseInternalServerError(gctx, err)
			return
		}

		req, err := http.NewRequestWithContext(gctx, gctx.Request.Method, url, bytes.NewReader(bodyBuffer))
		if err != nil {
			apiutil.ApiResponseInternalServerError(gctx, err)
			return
		}

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			apiutil.ApiResponseInternalServerError(gctx, err)
			return
		}
		var respBody Resp
		err = json.NewDecoder(resp.Body).Decode(&respBody)
		if err != nil {
			apiutil.ApiResponseInternalServerError(gctx, err)
			return
		}
		if resp.StatusCode != http.StatusOK {
			gctx.JSON(http.StatusInternalServerError, respBody)
			return
		}

		gctx.JSON(http.StatusOK, respBody)
	}
}
