//
//  Practicing Cassandra
//
//  Copyright © 2020. All rights reserved.
//

package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var starTime = time.Now()

// Ping will handle the ping endpoint
func Ping(c *gin.Context) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	startTime := starTime.In(loc)

	res := map[string]string{
		"start_time": startTime.Format("[02 January 2006] 15:04:05 MST"),
	}
	c.JSON(http.StatusOK, res)
}
