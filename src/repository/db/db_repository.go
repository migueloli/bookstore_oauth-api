package db

import (
	"github.com/gocql/gocql"
	"github.com/migueloli/bookstore_oauth-api/src/clients/cassandra"
	"github.com/migueloli/bookstore_oauth-api/src/domais/accesstoken"
	"github.com/migueloli/bookstore_oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_token WHERE access_token = ?;"
	queryCreateAccessToken = "INSERT INTO access_token(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_token SET expires = ? WHERE access_token = ?;"
)

// Repository is the controller for the database repository
type Repository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
	Create(accesstoken.AccessToken) *errors.RestErr
	UpdateExpirationTime(accesstoken.AccessToken) *errors.RestErr
}

type repository struct{}

// NewRepository creates and returns a new repository.
func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetByID(accessToken string) (*accesstoken.AccessToken, *errors.RestErr) {
	result := accesstoken.AccessToken{}
	if err := cassandra.GetSession().Query(queryGetAccessToken, accessToken).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError(err.Error())
		}

		return nil, errors.NewInternalServerError(err.Error())
	}

	return nil, nil
}

func (r *repository) Create(at accesstoken.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(
		queryCreateAccessToken,
		at.AccessToken,
		at.UserID,
		at.ClientID,
		at.Expires,
	).Exec(); err != nil {
		return errors.NewNotFoundError(err.Error())
	}

	return nil
}

func (r *repository) UpdateExpirationTime(at accesstoken.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(
		queryUpdateExpires,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return errors.NewNotFoundError(err.Error())
	}

	return nil
}
