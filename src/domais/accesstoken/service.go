package accesstoken

import (
	"strings"

	"github.com/migueloli/bookstore_oauth-api/src/utils/errors"
)

// Service is the base interface for the application services.
type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

// Repository is the base interface to be used to process the repository calls.
type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
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

func (s service) GetByID(accessTokenID string) (*AccessToken, *errors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.NewBadRequestError("Invalid access token.")
	}

	accessToken, err := s.repository.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (s service) Create(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.repository.Create(at)
}

func (s service) UpdateExpirationTime(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.repository.UpdateExpirationTime(at)
}
