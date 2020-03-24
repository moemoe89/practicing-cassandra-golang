//
//  Practicing Cassandra
//
//  Copyright © 2020. All rights reserved.
//

package form

// UserForm represent the user request model
type UserForm struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

// Validate represent the validation method from UserForm
func (v *UserForm) Validate() []string {
	errs := []string{}
	if len(v.Name) < 1 {
		errs = append(errs, "Name can't be empty")
	}

	if len(v.Gender) < 1 {
		errs = append(errs, "Gender can't be empty")
	}

	if v.Age < 1 {
		errs = append(errs, "Age can't be empty")
	}

	return errs
}
