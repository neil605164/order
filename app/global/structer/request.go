package structer

import "time"

type CreateReq struct {
	Username             string    `json:"username" validate:"required" example:"neil"`
	Email                string    `json:"email" validate:"required,email" example:"neil605164@gmail.com"`
	Password             string    `json:"password" validate:"required,min=6,max=20,eqfield=PasswordConfirmation" example:"qwer1234"`
	PasswordConfirmation string    `json:"passwordConfirmation" validate:"required,min=6,max=20" example:"qwer1234"`
	Birthday             time.Time `json:"birthday" validate:"required" example:"2022-07-24T00:00:00Z"`
	MemberNo             string    `json:"-"`
	Pwd                  string    `json:"-"`
}
