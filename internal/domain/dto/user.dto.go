package dto

import "github.com/bagusyanuar/go-erp/internal/domain/entity"

type UserDTO struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func ToUser(entity *entity.User) *UserDTO {
	return &UserDTO{
		ID:    entity.ID.String(),
		Email: entity.Email,
	}
}

func ToUsers(entities []entity.User) []UserDTO {
	users := make([]UserDTO, 0)
	for _, entity := range entities {
		e := *ToUser(&entity)
		users = append(users, e)
	}
	return users
}
