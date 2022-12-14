package models

import "time"

type Transaction struct {
	ID        int64          `json:"id"`
	UserID    int            `json:"user_id"`
	User      UserProfile    `json:"user"`
	Status    string         `json:"status"`
	Total     int            `json:"total"`
	Carts     []CartResponse `json:"carts"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
}

type TransactionResponse struct {
	ID     int64 `json:"id"`
	UserID int   `json:"user_id"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
