package domain

import (
	"time"
)

type Campaign struct {
	ID                   string    `db:"id"`
	CommerceID           string    `db:"commerce_id"`
	BranchID             string    `db:"branch_id"`
	CampaignName         string    `db:"campaign_name"`
	CampaignType         string    `db:"campaign_type"` //points or casback
	CampaignMultiplier   float64   `db:"campaign_multiplier"`
	CampaignPercentBonus float64   `db:"campaign_percent_bonus"`
	StartDate            time.Time `db:"start_date"`
	EndDate              time.Time `db:"end_date"`
}
