package transaction

import "time"

type Transactions struct {
	ID int
	CampaignID int
	UserID int
	Amount int
	Status string
	Code string
	CreatedAt time.Time
	UpdatedAt time.Time
}
