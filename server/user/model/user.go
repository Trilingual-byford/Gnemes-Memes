package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SexType int8

const (
	Man   SexType = 1
	MALE  SexType = 0
	Other SexType = 2
)

type User struct {
	ID                          primitive.ObjectID `json:"id" bson:"_id"`
	UserName                    string             `bson:"userName" json:"userName"`
	Email                       string             `bson:"email" json:"email"`
	Avatar                      string             `bson:"avatar" json:"avatar"`
	HashedPassword              []byte             `bson:"passWord" json:"passWord"`
	Roles                       []Role             `bson:"roles" json:"roles"`
	Sex                         SexType            `bson:"sex" json:"sex"`
	RegistryDate                bool               `bson:"registryDate" json:"registryDate"`
	LastLoginTime               bool               `bson:"lastLoginTime" json:"lastLoginTime"`
	IsCredentialsNonExpiredFlag bool               `bson:"isCredentialsNonExpiredFlag" json:"isCredentialsNonExpiredFlag"`
	IsActivating                bool               `bson:"isActivating" json:"isActivating"`
	LikedCollections            []string           `bson:"likedCollections" json:"likedCollections"`
	SavedCollection             []string           `bson:"savedCollection" json:"savedCollection"`
}
