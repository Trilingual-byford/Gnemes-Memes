package model

import "time"

type SexType int8

const (
	Man   SexType = 1
	MALE  SexType = 0
	Other SexType = 2
)

type User struct {
	UserName                    string    `bson:"userName" json:"userName"`
	Email                       string    `bson:"email" json:"email"`
	Avatar                      string    `bson:"avatar" json:"avatar"`
	HashedPassword              string    `bson:"passWord" json:"passWord"`
	Roles                       []Role    `bson:"roles" json:"roles"`
	Sex                         SexType   `bson:"sex" json:"sex"`
	RegistryDate                time.Time `bson:"registryDate" json:"registryDate"`
	LastLoginTime               time.Time `bson:"lastLoginTime" json:"lastLoginTime"`
	IsCredentialsNonExpiredFlag bool      `bson:"isCredentialsNonExpiredFlag" json:"isCredentialsNonExpiredFlag"`
	IsActivating                bool      `bson:"isActivating" json:"isActivating"`
	LikedCollections            []string  `bson:"likedCollections" json:"likedCollections"`
	SavedCollection             []string  `bson:"savedCollection" json:"savedCollection"`
}
