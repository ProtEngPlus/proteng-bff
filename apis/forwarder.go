package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"

	"github.com/protengplus/proteng-bff/models"
	"github.com/protengplus/proteng-bff/utils/apiutil"
)

// Forward the request to another service with the common response format
func Forward(baseUrl string) func(c *gin.Context) {
	return ForwardStrict[models.HttpResponse, interface{}](baseUrl)
}

func ForwardAddParam(baseUrl string, params map[string]interface{}) func(c *gin.Context) {
	return func(gctx *gin.Context) {
		// Construct URL with query parameters
		u, err := url.Parse(baseUrl)
		if err != nil {
			apiutil.ApiResponseInternalServerError(gctx, err)
			return
		}

		q := u.Query()
		for key, value := range params {
			q.Set(key, fmt.Sprintf("%v", value))
		}
		u.RawQuery = q.Encode()

		// Call ForwardStrict with the modified URL
		ForwardStrict[models.HttpResponse, interface{}](u.String())(gctx)
	}
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
