package user

import (
	"shoeshop-backend/src/interfaces/http/context"
	"shoeshop-backend/src/shared/constant"
)

func (s *service) UpdateUser(ctx *context.ApplicationContext, userInput UserUpdate) (resp *UserResponse, err error) {
	user, err := s.uRepo.GetById(ctx, userInput.ID)
	if err != nil {
		return nil, err
	}

	userInput.ToUserDomain(ctx, user)
	err = s.uRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	resp = entityToUserResponse(user)
	return
}

func (s *service) RegisterUser(ctx *context.ApplicationContext, userInput UserRegister) (resp *UserResponse, err error) {
	_, err = s.uRepo.GetByEmail(ctx, userInput.Email)
	if err == nil {
		return nil, constant.ErrorDataAlreadyExist
	} else if err != constant.ErrorDataNotFound {
		return nil, err
	}

	createUser := userInput.ToUserDomain(ctx)
	err = s.uRepo.Save(ctx, createUser)
	if err != nil {
		return nil, err
	}

	resp = entityToUserResponse(createUser)
	return
}

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

func (s *service) GetById(ctx *context.ApplicationContext, id string) (resp *UserResponse, err error) {
	user, err := s.uRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
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
