package domain

type Commerce struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	PointsXbuy  int64  `db:"points_x_buy"`  // points rewarded for purchases in stores
	ValueXPoint int64  `db:"value_x_point"` // how much money must be spent to obtain the points offered by the commerce
}
