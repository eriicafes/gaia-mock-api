package models

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/eriicafes/filedb"
)

const UserResource = "users"

type User struct {
	ID        filedb.ID `json:"id"`
	Name      string    `json:"name"`
	AccountID string    `json:"accountId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserQuery struct {
	ID        int
	AccountID string
	Name      string
}

func matchUserQuery(user User, query *UserQuery) bool {
	if query == nil {
		return true
	}

	matchID := user.ID == filedb.ID(query.ID)
	if query.ID == 0 {
		matchID = true
	}

	matchAccountID := user.AccountID == query.AccountID
	if query.AccountID == "" {
		matchAccountID = true
	}

	matchName := user.Name == query.Name
	if query.Name == "" {
		matchName = true
	}

	return matchID && matchAccountID && matchName
}

func generateUserId() string {
	rand.Seed(time.Now().UnixNano())

	alphabets := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	bytes := make([]rune, 21)

	for i := range bytes {
		bytes[i] = alphabets[rand.Intn(len(alphabets))]
	}

	return fmt.Sprintf("ga-%x", string(bytes))
}

func (m *model) getUsers() []User {
	var users []User

	m.db.Get(UserResource, &users)

	return users
}

func (m *model) setUsers(users []User) {
	data := make([]interface{}, 0, len(users))

	for _, user := range users {
		data = append(data, user)
	}

	m.db.Set(UserResource, data)
}

func (m *model) FindOneUser(query *UserQuery) (*User, error) {
	users := m.getUsers()

	for _, user := range users {
		match := matchUserQuery(user, query)

		if match {
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (m *model) FindManyUsers(query *UserQuery) []User {
	users := m.getUsers()
	var result []User

	for _, user := range users {
		match := matchUserQuery(user, query)

		if match {
			result = append(result, user)
		}
	}

	return result
}

func (m *model) CreateUser(user User) *User {
	users := m.getUsers()

	// override fields
	user.ID = filedb.ID(len(users) + 1)
	user.AccountID = generateUserId()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	users = append(users, user)

	m.setUsers(users)

	return &user
}

func (m *model) UpdateUser(query *UserQuery, updatedUser User) (*User, error) {
	users := m.getUsers()
	var newUsers []User

	var updated bool

	for _, user := range users {
		if updated {
			newUsers = append(newUsers, user)
			continue
		}

		match := matchUserQuery(user, query)

		if match {
			updated = true

			// override fields
			updatedUser.ID = user.ID
			updatedUser.AccountID = user.AccountID
			updatedUser.CreatedAt = user.CreatedAt
			updatedUser.UpdatedAt = time.Now()

			newUsers = append(newUsers, updatedUser)
			continue
		}

		newUsers = append(newUsers, user)
	}

	if !updated {
		return nil, errors.New("user not found")
	}

	m.setUsers(newUsers)
	return &updatedUser, nil
}

func (m *model) RemoveOneUser(query *UserQuery) error {
	users := m.getUsers()
	var newUsers []User

	var removed bool

	for _, user := range users {
		if removed {
			newUsers = append(newUsers, user)
			continue
		}

		match := matchUserQuery(user, query)

		if match {
			removed = true
			continue
		}

		newUsers = append(newUsers, user)
	}

	if !removed {
		return errors.New("user not found")
	}

	m.setUsers(newUsers)
	return nil
}

func (m *model) RemoveManyUsers(query *UserQuery) (int, error) {
	users := m.getUsers()
	var newUsers []User

	var count int

	for _, user := range users {
		match := matchUserQuery(user, query)

		if match {
			count++
			continue
		}

		newUsers = append(newUsers, user)
	}

	if count == 0 {
		return count, errors.New("users not found")
	}

	m.setUsers(newUsers)

	return count, nil
}
