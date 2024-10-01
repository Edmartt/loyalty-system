package dto

type CommerceDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PointsXbuy  int64  `json:"points_x_buy"`
	ValueXPoint int64  `json:"value_x_point"`
}
