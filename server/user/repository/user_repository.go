package repository

import (
	"context"
	"github.com/kataras/golog"
	"gnemes/common/config"
	"gnemes/common/model"
	"gnemes/common/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"time"
)

type UserRepository interface {
	Create(username, userId, password, email, avatar string, sex model.SexType) (model.User, error)
	GetByUserEmailAndPassword(email, password string) (model.User, bool)
	GetAllGnemesColletionsByUserEmail(email string) (model.User, bool)
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

func (m *mongoUserRepository) GetAllGnemesColletionsByUserEmail(email string) (model.User, bool) {
	panic("implement me")
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

func (m *mongoUserRepository) Create(username, userId, hashedPassword, email, avatar string, sex model.SexType) (model.User, error) {
	//user := model.User(username,email,avatar,hashedPassword,sex,time.Now(),time.Now(),true,true,nil,nil)
	roles := []model.Role{model.USER}
	user := model.User{username, email, userId, avatar, hashedPassword, roles, sex, time.Now(), time.Now(), true, true, nil, nil}
	result, err := m.userCollection.InsertOne(context.Background(), user)
	if err != nil {
		m.logger.Error("failed to save userInfo", err)
	} else {
		m.logger.Error("saved successfully", result)
	}
	return user, err
}

func (m *mongoUserRepository) GetByUserEmailAndPassword(userEmail, password string) (model.User, bool) {
	var users model.User
	err := m.userCollection.FindOne(context.Background(), bson.M{"email": userEmail}).Decode(&users)
	if err != nil {
		m.logger.Error("Failed to find Use from Db due to:", err)
		return users, false
	} else {
		isValidated := utils.ValidatePassword(password, []byte(users.HashedPassword))
		if isValidated {
			return users, true
		} else {
			return users, false
		}
	}

}

func (m *mongoUserRepository) GetAll() ([]model.User, error) {
	filter := bson.D{{}}
	cur, err := m.userCollection.Find(context.Background(), filter)
	var users []model.User
	for cur.Next(context.Background()) {
		user := model.User{}
		err := cur.Decode(&user)
		if err != nil {
			m.logger.Error("decode user error", err)
		} else {
			users = append(users, user)
		}
	}
	return users, err
}

//func ()  {
//
//}
