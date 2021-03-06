package model

import (
	"errors"
	"time"
)

type SexType string

const (
	Male   SexType = "male"
	FEMALE SexType = "female"
	Other  SexType = "other"
)

type User struct {
	UserName                    string    `bson:"userName" json:"userName"`
	Email                       string    `bson:"email" json:"email"`
	UserId                      string    `bson:"userId" json:"userId"`
	Avatar                      string    `bson:"avatar" json:"avatar"`
	HashedPassword              string    `bson:"passWord" json:"passWord"`
	Roles                       []Role    `bson:"roles" json:"roles"`
	Sex                         SexType   `bson:"sex" json:"sex"`
	RegistryDate                time.Time `bson:"registryDate" json:"registryDate"`
	LastLoginTime               time.Time `bson:"lastLoginTime" json:"lastLoginTime"`
	IsCredentialsNonExpiredFlag bool      `bson:"isCredentialsNonExpiredFlag" json:"isCredentialsNonExpiredFlag"`
	IsActivating                bool      `bson:"isActivating" json:"isActivating"`
	//NativeLanguage
	//Location
	//Country
	Collections []CollectedItem `bson:"collections" json:"collections"`
	Likes       []CollectedItem `bson:"likes" json:"likes"`
	Histories   []History       `bson:"histories" json:"histories"`
}
type CollectedItem struct {
	CollectedDate time.Time
}
type History struct {
	GnemesId  string    `bson:"gnemesId" json:"gnemesId"`
	CheckTime time.Time `bson:"registryDate" json:"registryDate"`
}

func GetSexTypeFromString(x string) (SexType, error) {
	var sexType SexType
	switch x {
	case string(Male):
		sexType = Male
		return sexType, nil
	case string(FEMALE):
		return sexType, nil
	case string(Other):
		return sexType, nil
	default:
		return sexType, errors.New("sex type is not exist")
	}
}
