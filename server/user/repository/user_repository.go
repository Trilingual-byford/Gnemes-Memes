package repository

import (
	"gnemes/user/model"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

type UserRepository interface {
	Create(username, password, email, avatar string, sex model.SexType) (model.User, error)
	GetByUsernameAndPassword(username, password string) (model.User, bool)
	GetAll() ([]model.User, error)
}

var (
	_ UserRepository = (*mongoUserRepository)(nil)
)

type mongoUserRepository struct {
	userCollection *mongo.Collection
	mu             sync.RWMutex
}

func NewMongoUserRepository() UserRepository {
	m := new(mongoUserRepository)
	return m
}

func (m *mongoUserRepository) Create(username, password, email, avatar string, sex model.SexType) (model.User, error) {
	panic("implement me")
}

func (m *mongoUserRepository) GetByUsernameAndPassword(username, password string) (model.User, bool) {
	panic("implement me")
}

func (m *mongoUserRepository) GetAll() ([]model.User, error) {
	panic("implement me")
}
