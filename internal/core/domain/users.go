package domain

import (
	"time"
)

type User struct {
	ID          string    `json:"id"`
	Fullname    string    `json:"fullname" validate:"required,"`
	PhoneNumber string    `json:"phone_number" validate:"required"`
	Role        string    `json:"role"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Status      int32     `json:"status"`
	RegionID    int32     `json:"region_id" validate:"required"`
	DistrictID  int32     `json:"district_id" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserAllResp struct {
	Users []User `json:"users"`
	Count int    `json:"count"`
}

type UserCreate struct {
	Fullname    string `json:"fullname,omitempty" validate:"required,"`
	PhoneNumber string `json:"phone_number,omitempty" validate:"required"`
	RegionID    int32  `json:"region_id,omitempty" validate:"required"`
	DistrictID  int32  `json:"district_id,omitempty" validate:"required"`
}
