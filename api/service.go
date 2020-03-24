//
//  Practicing Cassandra
//
//  Copyright © 2020. All rights reserved.
//

package api

import (
	"github.com/moemoe89/practicing-cassandra-golang/api/api_struct/form"
	"github.com/moemoe89/practicing-cassandra-golang/api/api_struct/model"

	"errors"
	"net/http"
	"time"

	"github.com/gocql/gocql"
	"github.com/sirupsen/logrus"
)

// Service represent the services
type Service interface {
	Create(user *form.UserForm) (*model.UserModel, int, error)
	Find() ([]*model.UserModel, int, error)
	FindByID(id string) (*model.UserModel, int, error)
	Update(user *form.UserForm, id string) (*model.UserModel, int, error)
	Delete(id string) (int, error)
}

type implService struct {
	log        *logrus.Entry
	repository Repository
}

// NewService will create an object that represent the Service interface
func NewService(log *logrus.Entry, r Repository) Service {
	return &implService{log: log, repository: r}
}

func (s *implService) Create(req *form.UserForm) (*model.UserModel, int, error) {
	user := &model.UserModel{
		ID:        gocql.TimeUUID(),
		Name:      req.Name,
		Gender:    req.Gender,
		Age:       req.Age,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	user, err := s.repository.Create(user)
	if err != nil {
		s.log.Errorf("can't create data: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return user, 0, nil
}

func (s *implService) Find() ([]*model.UserModel, int, error) {
	users, err := s.repository.Find()
	if err != nil {
		s.log.Errorf("can't get data: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return users, 0, nil
}

func (s *implService) FindByID(id string) (*model.UserModel, int, error) {
	user, err := s.repository.FindByID(id)
	if err == gocql.ErrNotFound {
		return nil, http.StatusNotFound, errors.New("Data not found")
	}

	if err != nil {
		s.log.Errorf("can't get detail data: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return user, 0, nil
}

func (s *implService) Update(req *form.UserForm, id string) (*model.UserModel, int, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		s.log.Errorf("can't get detail data: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	user.Name      = req.Name
	user.Gender    = req.Gender
	user.Age       = req.Age
	user.UpdatedAt = time.Now().UTC()

	user, err = s.repository.Update(user, id)
	if err != nil {
		s.log.Errorf("can't update data: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return user, 0, nil
}

func (s *implService) Delete(key string) (int, error) {
	err := s.repository.Delete(key)
	if err != nil {
		s.log.Errorf("can't delete data: %s", err.Error())
		return http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return 0, nil
}