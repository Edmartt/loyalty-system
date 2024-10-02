package dto

type CampaignDTO struct {
	ID                   string  `json:"id"`
	CommerceID           string  `json:"commerce_id"`
	BranchID             string  `json:"branch_id"`
	CampaignName         string  `json:"campaign_name"`
	CampaignType         string  `json:"campaign_type"`
	CampaignMultiplier   float64 `json:"campaign_multiplier"`
	CampaignPercentBonus float64 `json:"campaign_percent_bonus"`
	StartDate            string  `json:"start_date"`
	EndDate              string  `json:"end_date"`
}
