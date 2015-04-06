package client

import "github.com/dbongo/hackapp/model"

// UserService ...
type UserService struct {
	*Client
}

// GetCurrent GET /api/user
func (s *UserService) GetCurrent() (*model.User, error) {
	var user = model.User{}
	var err = s.run("GET", "/api/user", nil, &user)
	return &user, err
}
