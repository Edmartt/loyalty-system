package domain

import "time"

type BranchCampaign struct {
	CampaignID           string    `json:"campaign_id" db:"campaign_id"`
	CampaignName         string    `json:"campaign_name" db:"campaign_name"`
	CampaignMultiplier   int       `json:"campaign_multiplier" db:"campaign_multiplier"`
	CampaignPercentBonus int       `json:"campaign_percent_bonus" db:"campaign_percent_bonus"`
	StartDate            time.Time `json:"start_date" db:"start_date"`
	EndDate              time.Time `json:"end_date" db:"end_date"`
	BranchID             string    `json:"branch_id" db:"branch_id"`
	CommerceID           string    `json:"commerce_id" db:"commerce_id"`
	CommerceName         string    `json:"commerce_name" db:"commerce_name"`
	Address              string    `json:"address" db:"address"`
}
