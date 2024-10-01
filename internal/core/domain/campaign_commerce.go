package domain

import "time"

type CommerceCampaign struct {
	CampaignID   string `json:"campaign_id" db:"campaign_id"`
	CampaignName string `json:"campaign_name" db:"campaign_name"`
	float64
	CampaignMultiplier   float64   `json:"campaign_multiplier" db:"campaign_multiplier"`
	CampaignPercentBonus float64   `json:"campaign_percent_bonus" db:"campaign_percent_bonus"`
	StartDate            time.Time `json:"start_date" db:"start_date"`
	EndDate              time.Time `json:"end_date" db:"end_date"`
	CommerceID           string    `json:"commerce_id" db:"commerce_id"`
	CommerceName         string    `json:"commerce_name" db:"commerce_name"`
	PointsXBuy           float64   `json:"points_x_buy" db:"points_x_buy"`
	ValueXPoint          float64   `json:"value_x_point" db:"value_x_point"`
}
