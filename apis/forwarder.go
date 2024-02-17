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

// ForwardNoStrict is Forward but the request and response body can be anything
func ForwardNoStrict(baseUrl string, params ...map[string]interface{}) func(c *gin.Context) {
	return ForwardStrict[interface{}, interface{}](baseUrl, params...)
}

// Forward the request to another service with the common response format
func Forward(baseUrl string, params ...map[string]interface{}) func(c *gin.Context) {
	return ForwardStrict[models.HttpResponse, interface{}](baseUrl, params...)
}

// ForwardStrict will have to specify the request and response struct
func ForwardStrict[Resp any, Req any](baseUrl string, params ...map[string]interface{}) func(c *gin.Context) {
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

		// Construct URL with query parameters
		u, err := url.Parse(baseUrl)
		if err != nil {
			apiutil.ApiResponseInternalServerError(gctx, err)
			return
		}

		if len(params) > 0 {
			q := u.Query()
			for key, value := range params[0] {
				// Convert the value to string before setting it in the query
				q.Set(key, fmt.Sprintf("%v", value))
			}
			u.RawQuery = q.Encode()
		}

		fmt.Println(u)

		req, err := http.NewRequestWithContext(gctx, gctx.Request.Method, u.String(), bytes.NewReader(bodyBuffer))
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
		defer resp.Body.Close()

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
