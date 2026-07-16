package user

import "ecommerce/domain"

type service struct {
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) *service {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) Create(user domain.User) (*domain.User, error) {
	usr, error := s.userRepo.Create(user)
	if error != nil {
		return nil, error
	}
	if usr == nil {
		return nil, nil
	}
	return usr, nil
}

func (s *service) Find(email, pass string) (*domain.User, error) {
	user, err := s.userRepo.Find(email, pass)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return user, nil
}
