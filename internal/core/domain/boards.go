package domain

import (
	"time"
)

type Board struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Icon            string    `json:"icon"`
	TotalAmount     int       `json:"total_amount"`
	AcceptStartDate time.Time `json:"accept_start_date"`
	AcceptEndDate   time.Time `json:"accept_end_date"`
	ReviewStartDate time.Time `json:"review_start_date"`
	ReviewEndDate   time.Time `json:"review_end_date"`
	VotingStartDate time.Time `json:"voting_start_date"`
	VotingEndDate   time.Time `json:"voting_end_date"`
}

type BoardAllResp struct {
	Boards []Board `json:"boards"`
	Count  int     `json:"count"`
}

type BoardCreate struct {
	Title           string    `json:"title,omitempty"`
	Icon            string    `json:"icon,omitempty"`
	TotalAmount     int       `json:"total_amount,omitempty"`
	AcceptStartDate time.Time `json:"accept_start_date,omitempty"`
	AcceptEndDate   time.Time `json:"accept_end_date,omitempty"`
	ReviewStartDate time.Time `json:"review_start_date,omitempty"`
	ReviewEndDate   time.Time `json:"review_end_date,omitempty"`
	VotingStartDate time.Time `json:"voting_start_date,omitempty"`
	VotingEndDate   time.Time `json:"voting_end_date,omitempty"`
}
