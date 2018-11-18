package vars

import "github.com/amstee/easy-cut/src/common"

type UserCreation struct {
	Connection string 					`json:"connection"`
	Email string						`json:"email"`
	Password string						`json:"password"`
	EVerified bool						`json:"email_verified"`
	VerifyEmail	bool					`json:"verify_email"`
	UserMetadata common.UserMetadata 	`json:"user_metadata"`
}

type UserUpdate struct {
	Email string						`json:"email,omitempty"`
	UserMetadata common.UserMetadata	`json:"user_metadata,omitempty"`
}