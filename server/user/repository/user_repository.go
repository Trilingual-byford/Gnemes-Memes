package repository

import (
	"github.com/kataras/golog"
	"gnemes/common/config"
	"gnemes/user/model"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"time"
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

func NewMongoUserRepository(logger *golog.Logger) UserRepository {
	m := new(mongoUserRepository)
	mongoClient, err := config.GnemesDB(config.USER, logger)
	if err != nil {
		logger.Error("")
	}
	collection := mongoClient.Database("gnemes").Collection("GnemesPost")
	m.userCollection = collection
	return m
}

func (m *mongoUserRepository) Create(username, hashedPassword, email, avatar string, sex model.SexType) (model.User, error) {
	//user := model.User(username,email,avatar,hashedPassword,sex,time.Now(),time.Now(),true,true,nil,nil)
	roles := []model.Role{model.USER}
	user := model.User{username, email, avatar, hashedPassword, roles, sex, time.Now(), time.Now(), true, true, nil, nil}

}

func (m *mongoUserRepository) GetByUsernameAndPassword(username, password string) (model.User, bool) {
	panic("implement me")
}

func (m *mongoUserRepository) GetAll() ([]model.User, error) {
	panic("implement me")
}
