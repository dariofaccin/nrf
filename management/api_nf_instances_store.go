// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0

/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package management

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/free5gc/http_wrapper"
	"github.com/free5gc/nrf/logger"
	"github.com/free5gc/nrf/producer"
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
)

// GetNFInstances - Retrieves a collection of NF Instances
func HTTPGetNFInstances(c *gin.Context) {
	req := http_wrapper.NewRequest(c.Request, nil)
	req.Query = c.Request.URL.Query()

	httpResponse := producer.HandleGetNFInstancesRequest(req)

	responseBody, err := openapi.Serialize(httpResponse.Body, "application/json")
	if err != nil {
		logger.ManagementLog.Warnln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(httpResponse.Status, "application/json", responseBody)
	}
}
