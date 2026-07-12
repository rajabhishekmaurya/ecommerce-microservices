package service

import (
	"github.com/google/uuid"

	"github.com/rajabhishekmaurya/ecommerce-microservices/user-service/internal/model"
)

var users []model.User

func GetUsers() []model.User {
	return users
}

func CreateUser(user model.User) model.User {

	user.ID = uuid.New().String()

	users = append(users, user)

	return user
}

func GetUser(id string) (model.User, bool) {

	for _, u := range users {

		if u.ID == id {
			return u, true
		}

	}

	return model.User{}, false
}

func UpdateUser(id string, user model.User) bool {

	for i := range users {

		if users[i].ID == id {

			user.ID = id

			users[i] = user

			return true

		}

	}

	return false
}

func DeleteUser(id string) bool {

	for i := range users {

		if users[i].ID == id {

			users = append(users[:i], users[i+1:]...)

			return true

		}

	}

	return false
}
