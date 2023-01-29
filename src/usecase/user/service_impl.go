package user

import (
	"shoeshop-backend/src/interfaces/http/context"
	"shoeshop-backend/src/shared/constant"
)

func (s *service) LoginUser(ctx *context.ApplicationContext, email, password string) (resp *UserResponse, err error) {
	user, err := s.uRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !CheckPassword(password, user.Password) {
		return nil, constant.ErrorInvalidPassword
	}

	resp = entityToUserResponse(user)
	return
}

func (s *service) GetByEmail(ctx *context.ApplicationContext, email string) (resp *UserResponse, err error) {
	user, err := s.uRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	resp = entityToUserResponse(user)
	return
}

func (s *service) Migrate() {
	s.uRepo.AutoMigrate()
}
