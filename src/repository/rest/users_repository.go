package rest

import (
	"encoding/json"
	"time"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/migueloli/bookstore_oauth-api/src/domais/users"
	"github.com/migueloli/bookstore_oauth-api/src/utils/errors"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "localhost:8080",
		Timeout: 100 * time.Millisecond,
	}
)

// UserRepository is the controller for the rest repository
type UserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type userRepository struct{}

// NewRepository creates and returns a new repository.
func NewRepository() UserRepository {
	return &userRepository{}
}

func (r userRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)

	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("Invalid REST client response when trying to login with user.")
	}

	if response.StatusCode > 299 {
		restErr := errors.RestErr{}
		err := json.Unmarshal(response.Bytes(), &restErr)

		if err != nil {
			return nil, errors.NewInternalServerError("Invalid error interface. when trying to login with user.")
		}

		return nil, &restErr
	}

	user := users.User{}
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("Error when trying to unmarshall users response.")
	}

	return &user, nil
}
