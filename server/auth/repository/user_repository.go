package repository

import (
	"context"
	"github.com/kataras/golog"
	"gnemes/common/config"
	"gnemes/common/model"
	"go.mongodb.org/mongo-driver/bson"
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
	logger         *golog.Logger
}

func NewMongoUserRepository(logger *golog.Logger) UserRepository {
	m := new(mongoUserRepository)
	mongoClient, err := config.GetGnemesDBClient(config.USER, logger)
	m.logger = logger
	if err != nil {
		logger.Error("failed to init mongo Repository for User", err)
	}
	collection := mongoClient.Database("gnemes").Collection("GnemesUser")
	m.userCollection = collection
	return m
}

func (m *mongoUserRepository) Create(username, hashedPassword, email, avatar string, sex model.SexType) (model.User, error) {
	//auth := model.User(username,email,avatar,hashedPassword,sex,time.Now(),time.Now(),true,true,nil,nil)
	roles := []model.Role{model.USER}
	user := model.User{username, email, avatar, hashedPassword, roles, sex, time.Now(), time.Now(), true, true, nil, nil}
	result, err := m.userCollection.InsertOne(context.Background(), user)
	if err != nil {
		m.logger.Error("failed to save userInfo", err)
	} else {
		m.logger.Error("saved successfully", result)
	}
	return user, err
}

func (m *mongoUserRepository) GetByUsernameAndPassword(username, password string) (model.User, bool) {
	panic("implement me")
}

func (m *mongoUserRepository) GetAll() ([]model.User, error) {
	filter := bson.D{{}}
	cur, err := m.userCollection.Find(context.Background(), filter)
	var users []model.User
	for cur.Next(context.Background()) {
		user := model.User{}
		err := cur.Decode(&user)
		if err != nil {
			m.logger.Error("decode auth error", err)
		} else {
			users = append(users, user)
		}
	}
	return users, err
}
