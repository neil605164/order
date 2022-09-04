package models

import "time"

// UserReview 使用者審核資訊
type UserReview struct {
	Id
	UserID         uint64     `json:"userId" gorm:"comment:會員ID;NOT NULL;index:user_id"`
	VerifyType     string     `json:"verify_type" gorm:"type:varchar(30);comment:驗證類型;NOT NULL;index:verify_type"`
	Status         string     `json:"status" gorm:"type:varchar(30);comment:驗證狀態;NOT NULL;default:'not_verified';index:status"`
	SubmissionedAt *time.Time `json:"submissionedAt" gorm:"type:datetime;comment:送審時間;NULL"`
	ReviewedAt     *time.Time `json:"reviewedAt" gorm:"type:datetime;comment:審核時間;NULL"`
	CreatedAt      time.Time  `json:"created_at" gorm:"type:datetime;autoCreateTime;comment:建立時間;default:CURRENT_TIMESTAMP"`
	User           *User      `json:"user" gorm:"foreignKey:UserID"`
}
