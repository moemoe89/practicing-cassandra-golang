//
//  Practicing Cassandra
//
//  Copyright © 2020. All rights reserved.
//

package model

import (
	"time"

	"github.com/gocql/gocql"
)

// UserModel represent the user model
type UserModel struct {
	ID        gocql.UUID `json:"id" cql:"_id"`
	Name      string     `json:"name" cql:"name"`
	Gender    string     `json:"gender" cql:"gender"`
	Age       int        `json:"age" cql:"age"`
	CreatedAt time.Time  `json:"-" cql:"created_at"`
	UpdatedAt time.Time  `json:"-" cql:"updated_at"`
}