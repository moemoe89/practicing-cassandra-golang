//
//  Practicing Cassandra
//
//  Copyright © 2020. All rights reserved.
//

package api

import (
	"github.com/moemoe89/practicing-cassandra-golang/api/api_struct/model"

	"github.com/gocql/gocql"
)

// Repository represent the repositories
type Repository interface {
	Create(user *model.UserModel) (*model.UserModel, error)
	Find() ([]*model.UserModel, error)
	FindByID(id string) (*model.UserModel, error)
	Update(user *model.UserModel, id string) (*model.UserModel, error)
	Delete(id string) error
}

type cassandraRepository struct {
	Client *gocql.Session
}

// NewCassandraRepository will create an object that represent the Repository interface
func NewCassandraRepository(sess *gocql.Session) Repository {
	return &cassandraRepository{sess}
}

func (c *cassandraRepository) Create(user *model.UserModel) (*model.UserModel, error) {
	err := c.Client.Query(`INSERT INTO users (id, name, gender, age, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`, user.ID, user.Name, user.Gender, user.Age, user.CreatedAt, user.UpdatedAt).Exec()
	return user, err
}

func (c *cassandraRepository) Find() ([]*model.UserModel, error) {
	users := []*model.UserModel{}
	user := &model.UserModel{}
	iter := c.Client.Query("SELECT id, name, gender, age FROM users").Iter()
	for iter.Scan(&user.ID, &user.Name, &user.Gender, &user.Age) {
		users = append(users, user)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}
	return users, nil
}

func (c *cassandraRepository) FindByID(id string) (*model.UserModel, error) {
	uuid, _ := gocql.ParseUUID(id)
	user := &model.UserModel{}
	err := c.Client.Query(`SELECT id, name, gender, age FROM users WHERE id = ? LIMIT 1`,
		uuid).Consistency(gocql.One).Scan(&user.ID, &user.Name, &user.Gender, &user.Age)
	return user, err
}

func (c *cassandraRepository) Update(user *model.UserModel, id string) (*model.UserModel, error) {
	uuid, _ := gocql.ParseUUID(id)
	err := c.Client.Query(`UPDATE users SET name = ?, gender = ?, age = ?, updated_at = ? WHERE id = ?`, user.Name, user.Gender, user.Age, user.UpdatedAt, uuid).Exec()
	return user, err
}

func (c *cassandraRepository) Delete(id string) error {
	uuid, _ := gocql.ParseUUID(id)
	err := c.Client.Query(`DELETE FROM users WHERE id = ?`, uuid).Exec()
	return err
}
