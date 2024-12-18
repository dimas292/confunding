package transaction

import (
	"confunding/campaign"
	"confunding/user"
	"time"
)

type Transactions struct {
	ID int
	CampaignID int
	UserID int
	Amount int
	Status string
	Code string
	Campaign campaign.Campaign
	User user.User
	CreatedAt time.Time
	UpdatedAt time.Time
}
