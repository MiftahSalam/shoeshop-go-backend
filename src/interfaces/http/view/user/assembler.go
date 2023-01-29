package user

import "shoeshop-backend/src/usecase/user"

func toUserResponse(user *user.UserResponse) *User {
	return &User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		IsAdmin:   user.IsAdmin,
		CreatedAt: user.CreatedAt,
	}
}
