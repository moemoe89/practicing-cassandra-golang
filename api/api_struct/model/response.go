//
//  Practicing Cassandra
//
//  Copyright © 2020. All rights reserved.
//

package model

// NewGenericResponse will create an object that represent the GenericResponse struct
func NewGenericResponse(stsCd, isError int, messages []string, data interface{}) *GenericResponse {

	return &GenericResponse{
		Status:   stsCd,
		Success:  isError == 0,
		Messages: messages,
		Data:     data,
	}
}

// GenericResponse represent the generic response API
type GenericResponse struct {
	Status   int         `json:"status"`
	Success  bool        `json:"success"`
	Messages []string    `json:"messages"`
	Data     interface{} `json:"data"`
}
