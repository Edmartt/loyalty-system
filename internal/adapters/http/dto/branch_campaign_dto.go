package dto

type BranchCampaign struct {
	CampaignID           string  `json:"campaign_id" db:"campaign_id"`
	CampaignName         string  `json:"campaign_name" db:"campaign_name"`
	CampaignType         string  `json:"campaign_type" db:"campaign_type"`
	CampaignMultiplier   float64 `json:"campaign_multiplier" db:"campaign_multiplier"`
	CampaignPercentBonus float64 `json:"campaign_percent_bonus" db:"campaign_percent_bonus"`
	PointsXBuy           float64 `json:"points_x_buy" db:"points_x_buy"`
	ValueXPoint          float64 `json:"value_x_point" db:"value_x_point"`
	StartDate            string  `json:"start_date" db:"start_date"`
	EndDate              string  `json:"end_date" db:"end_date"`
	BranchID             string  `json:"branch_id" db:"branch_id"`
	CommerceID           string  `json:"commerce_id" db:"commerce_id"`
	CommerceName         string  `json:"commerce_name" db:"commerce_name"`
	Address              string  `json:"address" db:"address"`
}
