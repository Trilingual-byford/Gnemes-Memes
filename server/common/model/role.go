package model

type Role string

const (
	Admin        Role = "admin"
	USER         Role = "auth"
	PREMIUM_USER Role = "PremiumUser"
)
