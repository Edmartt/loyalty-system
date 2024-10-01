package dto

import "time"

type CampaignDTO struct {
	ID                   string    `json:"id"`
	CommerceID           string    `json:"commerce_id"`
	BranchID             string    `json:"branch_id"`
	CampaignName         string    `json:"campaign_name"`
	CampaignType         string    `json:"campaign_type"`
	CampaignMultiplier   int32     `json:"campaign_multiplier"`
	CampaignPercentBonus int16     `json:"campaign_percent_bonus"`
	StartDate            time.Time `json:"start_date"`
	EndDate              time.Time `json:"end_date"`
}
