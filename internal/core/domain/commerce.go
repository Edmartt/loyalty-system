package domain

type Commerce struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	PointsXbuy  int64  `db:"points_x_buy"`
	ValueXPoint int64  `db:"value_x_point"`
}
