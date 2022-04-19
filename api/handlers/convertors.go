package handlers

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/domain"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/gen/models"
)

func asUserResponse(d *domain.User) *models.User {
	genModelUsr := &models.User{
		ID:        d.ID,
		Name:      swag.String(d.Name),
		Address:   d.Address,
		CreatedAt: strfmt.DateTime(d.CreatedAt),
	}
	return genModelUsr
}
func asErrorResponse(e *domain.Error) *models.Error {
	er := &models.Error{
		Code:    int64(e.Code),
		Message: swag.String(e.Message),
	}
	return er
}
