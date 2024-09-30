package domain

// Reward entitie handles rewards redeemed with leal points
type Reward struct {
	ID           string `db:"id"`
	Name         string `db:"name"`
	CommerceID   string `db:"commerce_id"`
	PointsRedeem int64  `db:"points_redeem"`
}
