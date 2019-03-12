// Package user provides abstraction layer to User core business logic
package user

// Service concrete type for user service
type Service struct {
	repo Repoer
	key  []byte
}

// InitUserService initializes user service
func InitUserService(repo Repoer) *Service {
	// TODO: get key from config object
	key := "secret"
	return &Service{
		repo,
		[]byte(key),
	}
}

// CreateUser implements contract CreateUser of User Servicer interface
func (s *Service) CreateUser(user *User) error {
	return nil
}

// Login implements contract GetUser of User Servicer interface
func (s *Service) Login(username, password string) (*Login, error) {
	return nil, nil
}
