package db

import (
	"github.com/migueloli/bookstore_oauth-api/src/domais/accesstoken"
	"github.com/migueloli/bookstore_oauth-api/src/utils/errors"
)

// Repository is the controller for the database repository
type Repository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
}

type repository struct {
}

// NewRepository creates and returns a new repository.
func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetByID(string) (*accesstoken.AccessToken, *errors.RestErr) {
	return nil, nil
}
