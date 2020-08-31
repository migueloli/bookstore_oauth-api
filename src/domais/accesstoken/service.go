package accesstoken

import "github.com/migueloli/bookstore_oauth-api/src/utils/errors"

// Service is the base interface for the application services.
type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

// NewService creates and returns a new service.
func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

// GetByID is a method to get a AccessToken by its ID.
func (s service) GetByID(accessTokenID string) (*AccessToken, *errors.RestErr) {
	accessToken, err := s.repository.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}
