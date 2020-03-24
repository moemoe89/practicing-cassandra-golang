//
//  Practicing Cassandra
//
//  Copyright © 2020. All rights reserved.
//

package http

import (
	"github.com/moemoe89/practicing-cassandra-golang/api/api_struct/form"
	"github.com/moemoe89/practicing-cassandra-golang/api/api_struct/model"
	ap "github.com/moemoe89/practicing-cassandra-golang/api"
	cons "github.com/moemoe89/practicing-cassandra-golang/constant"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ctrl struct represent the delivery for controller
type ctrl struct {
	log *logrus.Entry
	svc ap.Service
}

// NewCtrl will create an object that represent the ctrl struct
func NewCtrl(log *logrus.Entry, svc ap.Service) *ctrl {
	return &ctrl{log, svc}
}

func (ct *ctrl) Create(c *gin.Context) {
	req := &form.UserForm{}
	if err := c.BindJSON(&req); err != nil {
		ct.log.Errorf("can't get json body: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewGenericResponse(http.StatusBadRequest, cons.ERR, []string{"Oops! Something went wrong with your request"}, nil))
		return
	}

	errs := req.Validate()
	if len(errs) > 0 {
		c.JSON(http.StatusBadRequest, model.NewGenericResponse(http.StatusBadRequest, cons.ERR, errs, nil))
		return
	}

	user, status, err := ct.svc.Create(req)
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	c.JSON(http.StatusCreated, model.NewGenericResponse(http.StatusCreated, cons.OK, []string{"Data has been saved"}, user))
}

func (ct *ctrl) Find(c *gin.Context) {
	users, status, err := ct.svc.Find()
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	c.JSON(http.StatusOK, model.NewGenericResponse(http.StatusOK, cons.OK, []string{"Data has been retrieved"}, users))
}

func (ct *ctrl) FindByID(c *gin.Context) {
	id := c.Param("id")

	user, status, err := ct.svc.FindByID(id)
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	c.JSON(http.StatusOK, model.NewGenericResponse(http.StatusOK, cons.OK, []string{"Data has been retrieved"}, user))
}

func (ct *ctrl) Update(c *gin.Context) {
	id := c.Param("id")

	req := &form.UserForm{}
	if err := c.BindJSON(&req); err != nil {
		ct.log.Errorf("can't get json body: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewGenericResponse(http.StatusBadRequest, cons.ERR, []string{"Oops! Something went wrong with your request"}, nil))
		return
	}

	errs := req.Validate()
	if len(errs) > 0 {
		c.JSON(http.StatusBadRequest, model.NewGenericResponse(http.StatusBadRequest, cons.ERR, errs, nil))
		return
	}

	user, status, err := ct.svc.FindByID(id)
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	user, status, err = ct.svc.Update(req, id)
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	c.JSON(http.StatusOK, model.NewGenericResponse(http.StatusOK, cons.OK, []string{"Data has been updated"}, user))
}

func (ct *ctrl) Delete(c *gin.Context) {
	id := c.Param("id")

	status, err := ct.svc.Delete(id)
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	c.JSON(http.StatusOK, model.NewGenericResponse(http.StatusOK, cons.OK, []string{"Data has been deleted"}, nil))
}