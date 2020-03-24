//
//  Practicing Cassandra
//
//  Copyright © 2020. All rights reserved.
//

package routers

import (
	ap "github.com/moemoe89/practicing-cassandra-golang/api"
	h "github.com/moemoe89/practicing-cassandra-golang/api/delivery/http"
	mw "github.com/moemoe89/practicing-cassandra-golang/api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GetRouter will create a variable that represent the gin.Engine
func GetRouter(log *logrus.Entry, svc ap.Service) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(mw.CORS)
	r.GET("/", ap.Ping)
	r.GET("/ping", ap.Ping)

	ctrl := h.NewCtrl(log, svc)

	r.POST("/user", ctrl.Create)
	r.GET("/user", ctrl.Find)
	r.GET("/user/:id", ctrl.FindByID)
	r.PUT("/user/:id", ctrl.Update)
	r.DELETE("/user/:id", ctrl.Delete)

	return r
}
