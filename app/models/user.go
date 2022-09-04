package models

import "time"

// User 使用者資訊
type User struct {
	Id
	ReviewID uint64    `json:"review_id" gorm:"unsigned;comment:審核列表ID;index"`
	MemberNo string    `json:"member_no" gorm:"type:varchar(30) comment '會員編號';not null;unique"`
	Username string    `json:"username" gorm:"type:varchar(30) comment '會員暱稱';not null"`
	Email    string    `json:"email" gorm:"type:varchar(255);unique;comment:信箱;not null"`
	Password string    `json:"-" gorm:"type:varchar(255);comment:密碼;not null"`
	Birthday time.Time `json:"birthday" gorm:"comment:生日;not null"`
	BaseTime
	Review *UserReview `gorm:"foreignKey:UserID"`
}
