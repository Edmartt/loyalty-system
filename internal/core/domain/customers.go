package domain

// Customer is an entity handling commerce customers
type Customer struct {
	ID                string `db:"id"`
	UserDNI           string `db:"user_dni"`
	Name              string `db:"last_name"`
	LastName          string `db:"last_name"`
	Phone             string `db:"phone"`
	Email             string `db:"email"`
	PointsCollected   int64  `db:"points_collected"`
	CashbackCollected int64  `db:"cashback_collected"`
}
