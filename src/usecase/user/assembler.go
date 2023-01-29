package user

import "shoeshop-backend/src/domain/user"

func entityToUserResponse(entity *user.User) *UserResponse {
	return &UserResponse{
		ID:        entity.ID.String(),
		Name:      entity.Name,
		Email:     entity.Email,
		IsAdmin:   entity.IsAdmin,
		CreatedAt: &entity.CreatedDate,
	}
}
