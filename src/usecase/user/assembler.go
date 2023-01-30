package user

import (
	"shoeshop-backend/src/domain/user"
	"shoeshop-backend/src/interfaces/http/context"
	"time"

	"github.com/google/uuid"
)

func entityToUserResponse(entity *user.User) *UserResponse {
	return &UserResponse{
		ID:        entity.ID.String(),
		Name:      entity.Name,
		Email:     entity.Email,
		IsAdmin:   entity.IsAdmin,
		CreatedAt: &entity.CreatedDate,
	}
}

func (uR UserRegister) ToUserDomain(ctx *context.ApplicationContext) *user.User {
	hashedPassword, err := HashPassword(uR.Password)
	if err != nil {
		ctx.Logger.Error("failed create user domain. error hash password: " + err.Error())
		return nil
	}

	return &user.User{
		ID:          uuid.New(),
		Name:        uR.Name,
		Email:       uR.Email,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
		Password:    hashedPassword,
	}
}

func (uU UserUpdate) ToUserDomain(ctx *context.ApplicationContext, curUser *user.User) {
	if uU.Password != "" {
		hashedPassword, err := HashPassword(uU.Password)
		if err != nil {
			ctx.Logger.Error("failed create user domain. error hash password: " + err.Error())
		}

		curUser.Password = hashedPassword
	}

	if uU.Email != "" {
		curUser.Email = uU.Email
	}

	if uU.Name != "" {
		curUser.Name = uU.Name
	}
}
