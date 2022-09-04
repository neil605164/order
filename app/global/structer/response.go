package structer

import "time"

type UserList struct {
	ID        uint64    `json:"id" example:"1"`
	CreatedAt time.Time `json:"created_at" example:"2022-07-21T12:19:39-04:00"`
	MemberNo  string    `json:"member_no" example:"md0002"`
	Username  string    `json:"username" example:"ban"`
	Email     string    `json:"email" example:"ban@gmail.com"`
	Birthday  time.Time `json:"birthday" example:"1992-07-18T20:19:28-04:00"`
	Review    *Review   `json:"review"`
}

type Review struct {
	ID             uint64     `json:"id" example:"2"`
	VerifyType     string     `json:"verify_type" example:"id_card"`
	Status         string     `json:"status" example:"verifying"`
	SubmissionedAt *time.Time `json:"submissionedAt" example:"null"`
	ReviewedAt     *time.Time `json:"reviewedAt" example:"null"`
	CreatedAt      time.Time  `json:"created_at" example:"2022-07-22T12:19:50-04:00"`
}
